package consumers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/processors"
)

type DeployConsumer struct {
	log       *logan.Entry
	cfg       config.FactoryTracker
	ctx       context.Context
	processor ethereum.FactoryProcessor
}

func NewDeployConsumer(cfg config.Config, ctx context.Context) *DeployConsumer {
	return &DeployConsumer{
		log:       cfg.Log(),
		ctx:       ctx,
		cfg:       cfg.FactoryTracker(),
		processor: processors.NewFactoryProcessor(cfg, ctx),
	}
}

func (c *DeployConsumer) Consume(ch <-chan etherdata.ContractDeployedEvent) {
	running.WithBackOff(
		c.ctx,
		c.log,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					if err = c.processor.ProcessDeploy(event); err != nil {
						return errors.Wrap(err, "failed to process deploy event")
					}
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}
