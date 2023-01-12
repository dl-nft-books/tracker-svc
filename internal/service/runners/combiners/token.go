package combiners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

type TokenCombiner struct {
	tracker  *trackers.TokenTracker
	consumer *consumers.TokenConsumer

	logger  *logan.Entry
	ctx     context.Context
	address common.Address
}

func NewTokenCombiner(cfg config.Config, ctx context.Context, address common.Address) *TokenCombiner {
	return &TokenCombiner{
		tracker:  trackers.NewTokenTracker(cfg, ctx),
		consumer: consumers.NewTokenConsumer(cfg, ctx),
		logger:   cfg.Log(),
		ctx:      ctx,
		address:  address,
	}
}

func (c *TokenCombiner) ProduceAndConsumeMintEvents() {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing mint event consumer and producer for %s", c.address.String())
	go func() {
		ch := make(chan etherdata.SuccessfulMintEvent)
		go c.tracker.TrackMintEvents(c.address, ch)
		go c.consumer.ConsumeMintEvents(c.address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeTransferEvents() {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing transfer event consumer and producer for %s", c.address.String())
	go func() {
		ch := make(chan etherdata.TransferEvent)
		go c.tracker.TrackTransferEvents(c.address, ch)
		go c.consumer.ConsumeTransferEvents(c.address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeUpdateEvents() {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing update event consumer and producer for %s", c.address.String())
	go func() {
		ch := make(chan etherdata.UpdateEvent)
		go c.tracker.TrackUpdateEvents(c.address, ch)
		go c.consumer.ConsumeUpdateEvents(c.address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeVoucherUpdateEvents() {
	c.logger.Infof("Initializing voucher update event consumer and producer for %s", c.address.String())
	go func() {
		ch := make(chan etherdata.VoucherUpdateEvent)
		go c.tracker.TrackVoucherUpdateEvents(c.address, ch)
		//TODO consumer
		//go c.consumer.
	}()
}

func (c *TokenCombiner) ProduceAndConsumeAllEvents() {
	c.logger.Infof("Initializing all possible consumers and producers for %s", c.address.String())
	c.ProduceAndConsumeMintEvents()
	c.ProduceAndConsumeTransferEvents()
	c.ProduceAndConsumeUpdateEvents()
	c.ProduceAndConsumeVoucherUpdateEvents()
}
