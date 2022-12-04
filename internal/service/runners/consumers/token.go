package consumers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
)

type TokenConsumer struct {
	logger   *logan.Entry
	cfg      config.FactoryTracker
	ctx      context.Context
	database data.DB
}

func NewTokenConsumer(cfg config.Config, ctx context.Context) *TokenConsumer {
	return &TokenConsumer{
		logger:   cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.FactoryTracker(),
		database: postgres.NewDB(cfg.DB()),
	}
}

func (c *TokenConsumer) ConsumeMintEvents(address common.Address, ch <-chan etherdata.SuccessfulMintEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					// TODO: Add implementation
					c.logger.Infof("Hey, I found this nice little thing over here: %v", event)
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeTransferEvents(address common.Address, ch <-chan etherdata.TransferEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					// TODO: Add implementation
					c.logger.Infof("Hey, I found this nice little thing over here: %v", event)
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeUpdateEvents(address common.Address, ch <-chan etherdata.UpdateEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					// TODO: Add implementation
					c.logger.Infof("Hey, I found this nice little thing over here: %v", event)
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}
