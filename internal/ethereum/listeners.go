package ethereum

import (
	"context"
	"time"

	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/ethereum/go-ethereum/common"
)

type (
	MarketplaceListener interface {
		From(from uint64) MarketplaceListener
		To(to uint64) MarketplaceListener
		WithMaxDepth(maxDepth uint64) MarketplaceListener
		WithAddress(address common.Address) MarketplaceListener
		WithCtx(ctx context.Context) MarketplaceListener
		WithDelayBetweenIntervals(delay time.Duration) MarketplaceListener
	}

	TokenListener interface {
		From(from uint64) TokenListener
		To(to uint64) TokenListener
		WithMaxDepth(maxDepth uint64) TokenListener
		WithAddress(address common.Address) TokenListener
		WithCtx(ctx context.Context) TokenListener
		WithDelayBetweenIntervals(delay time.Duration) TokenListener

		WatchTokenSuccessfullyPurchasedEvents(ch chan<- etherdata.TokenSuccessfullyPurchasedEvent) error
		//WatchSuccessfulMintByNftEvents(ch chan<- etherdata.SuccessfullyMintedByNftEvent) error
	}
)
