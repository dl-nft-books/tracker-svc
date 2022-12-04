package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/consumers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/trackers"
)

func Run(cfg config.Config, ctx context.Context) {
	var (
		deployTracker  = trackers.NewDeployTracker(cfg, ctx)
		deployConsumer = consumers.NewDeployConsumer(cfg, ctx)
		ch             = make(chan etherdata.ContractDeployedEvent)
	)

	go deployTracker.Track(ch)
	go deployConsumer.Consume(ch)
}
