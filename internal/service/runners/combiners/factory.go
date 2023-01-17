package combiners

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

type FactoryCombiner struct {
	logger *logan.Entry

	tracker  *trackers.FactoryTracker
	consumer *consumers.FactoryConsumer
	mutex    *sync.RWMutex
}

func NewFactoryCombiner(cfg config.Config, ctx context.Context, network models.NetworkDetailedResponse) *FactoryCombiner {
	return &FactoryCombiner{
		logger: cfg.Log(),

		tracker:  trackers.NewFactoryTracker(cfg, ctx, network),
		consumer: consumers.NewFactoryConsumer(cfg, ctx),
	}
}

func (c *FactoryCombiner) ProduceAndConsumeDeployEvents(routinerChannel chan<- common.Address) {
	c.logger.Info("Running producer and consumer for a factory contract")
	go func() {
		internalChannel := make(chan etherdata.ContractDeployedEvent)
		go c.tracker.TrackDeployEvents(internalChannel)
		go c.consumer.ConsumeDeployedEvents(internalChannel, routinerChannel)
	}()
}
