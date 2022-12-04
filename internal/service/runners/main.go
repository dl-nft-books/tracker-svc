package runners

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
)

func Run(cfg config.Config, ctx context.Context) error {
	var (
		factoryCombiner = combiners.NewFactoryCombiner(cfg, ctx)
		tokenCombiner   = combiners.NewTokenCombiner(cfg, ctx)
		logger          = cfg.Log()
	)

	contracts, err := postgres.NewContractsQ(cfg.DB()).Select()
	if err != nil {
		return errors.Wrap(err, "failed to select contracts from the database")
	}
	for _, contract := range contracts {
		tokenCombiner.ProduceAndConsumeAllEvents(contract.Address())
		logger.Infof("Initialized all consumers and trackers for contract %d", contract.Id)
	}

	factoryCombiner.ProduceAndConsumeDeployEvents()
	return nil
}
