package combiners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

type TokenCombiner struct {
	tracker  *trackers.TokenTracker
	consumer *consumers.TokenConsumer
}

func NewTokenCombiner(cfg config.Config, ctx context.Context) *TokenCombiner {
	return &TokenCombiner{
		tracker:  trackers.NewTokenTracker(cfg, ctx),
		consumer: consumers.NewTokenConsumer(cfg, ctx),
	}
}

func (c *TokenCombiner) ProduceAndConsumeMintEvents(address common.Address) {
	// Running tracker (producer) and consumer with a channel joining them
	// TODO: Make as a runner.WithBackOff?
	go func() {
		ch := make(chan etherdata.SuccessfulMintEvent)
		go c.tracker.TrackMintEvents(address, ch)
		go c.consumer.ConsumeMintEvents(address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeTransferEvents(address common.Address) {
	// Running tracker (producer) and consumer with a channel joining them
	go func() {
		ch := make(chan etherdata.TransferEvent)
		go c.tracker.TrackTransferEvents(address, ch)
		go c.consumer.ConsumeTransferEvents(address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeUpdateEvents(address common.Address) {
	// Running tracker (producer) and consumer with a channel joining them
	go func() {
		ch := make(chan etherdata.UpdateEvent)
		go c.tracker.TrackUpdateEvents(address, ch)
		go c.consumer.ConsumeUpdateEvents(address, ch)
	}()
}
