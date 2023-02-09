package ethereum

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
)

type (
	FactoryListener interface {
		From(from uint64) FactoryListener
		To(to uint64) FactoryListener
		WithMaxDepth(maxDepth uint64) FactoryListener
		WithAddress(address common.Address) FactoryListener
		WithCtx(ctx context.Context) FactoryListener
		WithDelayBetweenIntervals(delay time.Duration) FactoryListener

		WatchContractCreatedEvents(ch chan<- etherdata.ContractDeployedEvent) error
	}

	TokenListener interface {
		From(from uint64) TokenListener
		To(to uint64) TokenListener
		WithMaxDepth(maxDepth uint64) TokenListener
		WithAddress(address common.Address) TokenListener
		WithCtx(ctx context.Context) TokenListener
		WithDelayBetweenIntervals(delay time.Duration) TokenListener

		WatchSuccessfulMintEvents(ch chan<- etherdata.SuccessfulMintEvent) error
		WatchSuccessfulMintByNftEvents(ch chan<- etherdata.SuccessfullyMintedByNftEvent) error
		WatchTransferEvents(ch chan<- etherdata.TransferEvent) error
		WatchUpdateEvents(ch chan<- etherdata.UpdateEvent) error
		WatchVoucherUpdateEvents(ch chan<- etherdata.VoucherUpdateEvent) error
	}
)
