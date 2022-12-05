package combiners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
)

type TokenRoutiner struct {
	combiner *TokenCombiner
	logger   *logan.Entry
	cfg      config.ContractTracker
	ctx      context.Context

	routinesNumber uint64
}

func NewTokenRoutiner(cfg config.Config, ctx context.Context) *TokenRoutiner {
	return &TokenRoutiner{
		logger:   cfg.Log(),
		cfg:      cfg.ContractTracker(),
		ctx:      ctx,
		combiner: NewTokenCombiner(cfg, ctx),
	}
}

// Watch listens to the newly deployed token contracts and runs
// consumer and producer for them while acting as a sort of 'routines' manager
func (r *TokenRoutiner) Watch(ch <-chan common.Address) {
	running.WithBackOff(
		r.ctx,
		r.logger,
		r.cfg.Name,
		func(ctx context.Context) error {
			for {
				select {
				case address := <-ch:
					r.combiner.ProduceAndConsumeAllEvents(address)

					r.routinesNumber++
					r.logger.Infof("Caught new token to run combiner for. Current number of tokens to process: %d\n", r.routinesNumber)
				}
			}
		},
		r.cfg.Runner.NormalPeriod,
		r.cfg.Runner.MinAbnormalPeriod,
		r.cfg.Runner.MaxAbnormalPeriod,
	)
}
