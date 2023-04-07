package runners

import (
	"context"
	"github.com/dl-nft-books/network-svc/connector/models"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/data/postgres"
	"github.com/dl-nft-books/tracker-svc/internal/service/runners/combiners"
	"github.com/dl-nft-books/tracker-svc/internal/service/runners/trackers"
	"github.com/dl-nft-books/tracker-svc/solidity/generated/contractsregistry"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"time"

	"github.com/dl-nft-books/tracker-svc/internal/config"
	"github.com/pkg/errors"
)

const (
	delayBetweenNetworkSVCCalls = time.Second
)

func Run(cfg config.Config, ctx context.Context) error {
	// Channel connecting factory deploy consumer and routiner
	var (
		deployedTokensCh = make(chan trackers.DeployedToken)
		networks         *models.NetworkDetailedListResponse
		err              error
	)
	networkConnector := cfg.NetworkConnector()
	for networks == nil {
		networks, err = networkConnector.GetNetworksDetailed()
		if err != nil {
			cfg.Log().Error(errors.Wrap(err, "failed to select networks from the database"))
		}
		time.Sleep(delayBetweenNetworkSVCCalls)
	}
	err = postgres.NewDB(cfg.DB()).KeyValue().Upsert(data.KeyValue{
		Key:   "stats-chain_id-total",
		Value: strconv.Itoa(len(networks.Data)),
	})
	if err != nil {
		cfg.Log().WithError(err).Debug("failed to update chains amount")
		return errors.Wrap(err, "failed to get marketplace contract")
	}
	// factoryCombiner would run producer and consumer for a factory contract
	// and, after consumer processes the event, consumer will
	// send address to the deployedTokensCh and ask routiner to run
	// producer and consumer for it
	for _, network := range networks.Data {
		contractsRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(network.FactoryAddress), network.RpcUrl)
		if err != nil {
			cfg.Log().WithError(err).Debug("failed to create contract registry")
			return errors.Wrap(err, "failed to get marketplace contract")
		}
		marketplaceContract, err := contractsRegistry.GetMarketplaceContract(nil)
		go func(address common.Address, network models.NetworkDetailedResponse) {
			deployedTokensCh <- trackers.DeployedToken{
				Address: marketplaceContract,
				Network: network,
			}
		}(marketplaceContract, network)
	}

	tokenRoutiner := combiners.NewTokenRoutiner(cfg, ctx)
	tokenRoutiner.Watch(deployedTokensCh)
	return nil
}
