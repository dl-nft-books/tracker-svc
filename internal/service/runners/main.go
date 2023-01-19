package runners

import (
	"context"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
)

const delayBetweenContractInsertions = time.Second
const delayBetweenCombinerCalls = time.Second
const delayBetweenNetworkSVCCalls = time.Second

func Run(cfg config.Config, ctx context.Context) error {
	// Channel connecting factory deploy consumer and routiner
	var deployedTokensCh = make(chan consumers.DeployedToken)

	networkConnector := cfg.NetworkConnector()
	var networks *models.NetworkDetailedListResponse
	var err error
	for networks == nil {
		networks, err = networkConnector.GetNetworksDetailed()
		if err != nil {
			cfg.Log().Error(errors.Wrap(err, "failed to select networks from the database"))
		}
		time.Sleep(delayBetweenNetworkSVCCalls)
	}

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
