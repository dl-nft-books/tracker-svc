package combiners

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

type TokenCombiner struct {
	tracker  *trackers.TokenTracker
	consumer *consumers.TokenConsumer

	logger *logan.Entry
	cfg    config.ContractTracker
	ctx    context.Context
}

func NewTokenCombiner(cfg config.Config, ctx context.Context) *TokenCombiner {
	return &TokenCombiner{
		tracker:  trackers.NewTokenTracker(cfg, ctx),
		consumer: consumers.NewTokenConsumer(cfg, ctx),
		logger:   cfg.Log(),

		cfg: cfg.ContractTracker(),
		ctx: ctx,
	}
}

func (c *TokenCombiner) ProcessNewTokenContracts(ch <-chan common.Address) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case addr := <-ch:
					c.ProduceAndConsumeAllEvents(addr)
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *TokenCombiner) ProduceAndConsumeMintEvents(address common.Address) {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	// TODO: Make as a runner.WithBackOff?
	c.logger.Infof("Initializing mint event consumer and producer for %s", address.String())
	go func() {
		ch := make(chan etherdata.SuccessfulMintEvent)
		go c.tracker.TrackMintEvents(address, ch)
		go c.consumer.ConsumeMintEvents(address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeTransferEvents(address common.Address) {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing transfer event consumer and producer for %s", address.String())
	go func() {
		ch := make(chan etherdata.TransferEvent)
		go c.tracker.TrackTransferEvents(address, ch)
		go c.consumer.ConsumeTransferEvents(address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeUpdateEvents(address common.Address) {
	// Running tracker (producer) and consumer with a combinersChannel joining them
	c.logger.Infof("Initializing consumer event consumer and producer for %s", address.String())
	go func() {
		ch := make(chan etherdata.UpdateEvent)
		go c.tracker.TrackUpdateEvents(address, ch)
		go c.consumer.ConsumeUpdateEvents(address, ch)
	}()
}

func (c *TokenCombiner) ProduceAndConsumeAllEvents(address common.Address) {
	c.logger.Infof("Initializing all possible consumers and producers for %s", address.String())
	c.ProduceAndConsumeMintEvents(address)
	c.ProduceAndConsumeTransferEvents(address)
	c.ProduceAndConsumeUpdateEvents(address)
}
