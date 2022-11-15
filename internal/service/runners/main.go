package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"time"
)

const delayBetweenTrackers = 5 * time.Second

type Runner func(ctx context.Context)

func initializeRunners(cfg config.Config) (runners []Runner) {
	runners = append(runners, NewFactoryTracker(cfg).Run)
	runners = append(runners, NewMintTracker(cfg).Run)
	runners = append(runners, NewTransferTracker(cfg).Run)

	return
}

func Run(cfg config.Config, ctx context.Context) {
	for _, runner := range initializeRunners(cfg) {
		go runner(ctx)
		time.Sleep(delayBetweenTrackers)
	}
}
