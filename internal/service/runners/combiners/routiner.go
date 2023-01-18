package combiners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
)

const tokensRoutinerSuffix = "-tokens"

type TokenRoutiner struct {
	logger    *logan.Entry
	runnerCfg config.Runner

	ctx            context.Context
	cfg            config.Config
	network        models.NetworkDetailedResponse
	routinesNumber uint64
}

func NewTokenRoutiner(cfg config.Config, ctx context.Context, network models.NetworkDetailedResponse) *TokenRoutiner {
	return &TokenRoutiner{
		logger:    cfg.Log(),
		runnerCfg: cfg.Routiner(),
		network:   network,
		cfg:       cfg,
		ctx:       ctx,
	}
}

// Watch listens to the newly deployed token contracts and runs
// consumer and producer for them while acting as a sort of 'routines' manager
func (r *TokenRoutiner) Watch(ch <-chan common.Address) {
	running.WithBackOff(
		r.ctx,
		r.logger,
		r.runnerCfg.Prefix+tokensRoutinerSuffix,
		func(ctx context.Context) error {
			for {
				select {
				case address := <-ch:
					r.routinesNumber++
					r.logger.Infof("Caught new token to run combiner for. Current number of tokens to process: %d\n", r.routinesNumber)

					NewTokenCombiner(r.cfg, r.ctx, address, r.network).ProduceAndConsumeAllEvents()
				}
			}
		},
		r.runnerCfg.Backoff.NormalPeriod,
		r.runnerCfg.Backoff.MinAbnormalPeriod,
		r.runnerCfg.Backoff.MaxAbnormalPeriod,
	)
}
