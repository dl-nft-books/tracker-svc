package combiners

import (
	"context"
	"github.com/dl-nft-books/tracker-svc/internal/config"
	"github.com/dl-nft-books/tracker-svc/internal/service/runners/trackers"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
)

const tokensRoutinerSuffix = "-tokens"

type TokenRoutiner struct {
	logger    *logan.Entry
	runnerCfg config.Runner

	ctx            context.Context
	cfg            config.Config
	routinesNumber uint64
}

func NewTokenRoutiner(cfg config.Config, ctx context.Context) *TokenRoutiner {
	return &TokenRoutiner{
		logger:    cfg.Log(),
		runnerCfg: cfg.Routiner(),
		cfg:       cfg,
		ctx:       ctx,
	}
}

// Watch listens to the newly deployed token contracts and runs
// consumer and producer for them while acting as a sort of 'routines' manager
func (r *TokenRoutiner) Watch(ch <-chan trackers.DeployedToken) {
	running.WithBackOff(
		r.ctx,
		r.logger,
		r.runnerCfg.Prefix+tokensRoutinerSuffix,
		func(ctx context.Context) error {
			for {
				select {
				case deployedToken := <-ch:
					r.routinesNumber++
					r.logger.Infof("Caught new token to run combiner for. Current number of tokens to process: %d\n", r.routinesNumber)
					NewTokenCombiner(r.cfg, r.ctx, deployedToken).ProduceAndConsumeAllEvents()
				}
			}
		},
		r.runnerCfg.Backoff.NormalPeriod,
		r.runnerCfg.Backoff.MinAbnormalPeriod,
		r.runnerCfg.Backoff.MaxAbnormalPeriod,
	)
}
