package combiners

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

type FactoryCombiner struct {
	tracker  *trackers.FactoryTracker
	consumer *consumers.FactoryConsumer
}

func NewFactoryCombiner(cfg config.Config, ctx context.Context) *FactoryCombiner {
	return &FactoryCombiner{
		tracker:  trackers.NewFactoryTracker(cfg, ctx),
		consumer: consumers.NewFactoryConsumer(cfg, ctx),
	}
}

func (c *FactoryCombiner) ProduceAndConsumeDeployEvents(combinersChannel chan<- common.Address) {
	go func() {
		internalChannel := make(chan etherdata.ContractDeployedEvent)
		go c.tracker.TrackDeployEvents(internalChannel)
		go c.consumer.ConsumeDeployedEvents(internalChannel, combinersChannel)
	}()
}
