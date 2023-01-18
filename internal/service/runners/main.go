package runners

import (
	"context"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"log"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
)

const delayBetweenContractInsertions = time.Second
const delayBetweenCombinerCalls = time.Second

func Run(cfg config.Config, ctx context.Context) error {
	// Channel connecting factory deploy consumer and routiner
	var deployedTokensCh = make(chan consumers.DeployedToken)

	networkConnector := cfg.NetworkConnector()
	networks, err := networkConnector.GetNetworksDetailed()
	if err != nil {
		return errors.Wrap(err, "failed to select networks from the database")
	}
	log.Println("chain_id", (*networks).Data[0].ChainId)

	// factoryCombiner would run producer and consumer for a factory contract
	// and, after consumer processes the event, consumer will
	// send address to the deployedTokensCh and ask routiner to run
	// producer and consumer for it

	for _, network := range networks.Data {
		fmt.Println("NETWOOORK", network.Name)
		contracts, err := postgres.NewContractsQ(cfg.DB()).FilterByChainId(network.ChainId).Select()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to select contracts from the database for network %v", network.Name))
		}
		for _, contract := range contracts {
			go func(contract data.Contract) {
				deployedTokensCh <- consumers.DeployedToken{
					Address: contract.Address(),
					Network: network,
				}
			}(contract)

			time.Sleep(delayBetweenContractInsertions)
		}
		factoryCombiner := combiners.NewFactoryCombiner(cfg, ctx, network)
		factoryCombiner.ProduceAndConsumeDeployEvents(deployedTokensCh)
		time.Sleep(delayBetweenCombinerCalls)
	}

	tokenRoutiner := combiners.NewTokenRoutiner(cfg, ctx)
	tokenRoutiner.Watch(deployedTokensCh)
	return nil
}
