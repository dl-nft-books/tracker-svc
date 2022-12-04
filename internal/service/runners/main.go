package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
	"time"
)

const delayBetweenTrackers = 5 * time.Second

type Runner func(ctx context.Context)

func initializeRunners(cfg config.Config, ctx context.Context) (runners []Runner) {
	runners = append(runners, trackers.NewFactoryTracker(cfg, ctx).Run)
	runners = append(runners, trackers.NewMintTracker(cfg).Run)
	runners = append(runners, trackers.NewTransferTracker(cfg).Run)

	return
}

func Run(cfg config.Config, ctx context.Context) {
	for _, runner := range initializeRunners(cfg, ctx) {
		go runner(ctx)
		time.Sleep(delayBetweenTrackers)
	}
}
