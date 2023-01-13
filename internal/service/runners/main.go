package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
)

const delayBetweenContractInsertions = time.Second
const delayBetweenCombinerCalls = time.Second

func Run(cfg config.Config, ctx context.Context) error {
	var (
		tokenRoutiner = combiners.NewTokenRoutiner(cfg, ctx)

		// Channel connecting factory deploy consumer and routiner
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

		time.Sleep(delayBetweenContractInsertions)
	}
	networkConnector := cfg.NetworkConnector()
	networks, err := networkConnector.GetNetworks()
	if err != nil {
		return errors.Wrap(err, "failed to select networks from the database")
	}

	// factoryCombiner would run producer and consumer for a factory contract
	// and, after consumer processes the event, consumer will
	// send address to the deployedTokensCh and ask routiner to run
	// producer and consumer for it
	tokenRoutiner.Watch(deployedTokensCh)
	for _, network := range networks.Data {
		factoryCombiner := combiners.NewFactoryCombiner(cfg, ctx, common.HexToAddress(network.FactoryAddress))
		factoryCombiner.ProduceAndConsumeDeployEvents(deployedTokensCh)

		time.Sleep(delayBetweenCombinerCalls)
	}

	return nil
}
