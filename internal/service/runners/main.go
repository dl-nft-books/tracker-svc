package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
)

func Run(cfg config.Config, ctx context.Context) error {
	var (
		factoryCombiner = combiners.NewFactoryCombiner(cfg, ctx)
		tokenRoutiner   = combiners.NewTokenRoutiner(cfg, ctx)

		// Special channel to link two combiners in order
		// to fetch newly deployed contracts and track them as well
		deployedTokensCh = make(chan common.Address)
	)

	contracts, err := postgres.NewContractsQ(cfg.DB()).Select()
	if err != nil {
		return errors.Wrap(err, "failed to select contracts from the database")
	}

	for _, contract := range contracts {
		go func(contract data.Contract) {
			deployedTokensCh <- contract.Address()
		}(contract)
	}

	// factoryCombiner would run producer and consumer for a factory contract
	// and, after consumer processes the event, consumer will
	// send address to the deployedTokensCh and ask routiner to run
	// producer and consumer for it
	factoryCombiner.ProduceAndConsumeDeployEvents(deployedTokensCh)
	tokenRoutiner.Watch(deployedTokensCh)

	return nil
}
