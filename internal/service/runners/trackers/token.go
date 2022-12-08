package trackers

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/listeners/token"
)

const (
	transferTrackerSuffix = "-token-transfer"
	mintTrackerSuffix     = "-token-mint"
	updateTrackerSuffix   = "-token-update"
)

type TokenTracker struct {
	log      *logan.Entry
	cfg      config.Trackers
	ctx      context.Context
	database data.DB
	listener ethereum.TokenListener
	mutex    *sync.RWMutex
}

func NewTokenTracker(cfg config.Config, ctx context.Context) *TokenTracker {
	mutex := new(sync.RWMutex)

	return &TokenTracker{
		log:      cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.Trackers(),
		database: postgres.NewDB(cfg.DB()),
		listener: token_listeners.NewTokenListener(cfg, ctx, mutex).
			WithMaxDepth(cfg.Trackers().MaxDepth),
		mutex: mutex,
	}
}

func (t *TokenTracker) TrackTransferEvents(address common.Address, ch chan<- etherdata.TransferEvent) {
	running.WithBackOff(
		t.ctx,
		t.log,
		t.cfg.Prefix+transferTrackerSuffix,
		func(ctx context.Context) error {
			startBlock := uint64(0)

			contract, err := t.database.Contracts().GetByAddress(address.String())
			if err != nil {
				return errors.Wrap(err, "failed to get contract by address")
			}
			if contract == nil {
				t.log.Warnf("The following contract is not contained in the database: %s", address.String())
				return t.listener.From(0).WithCtx(ctx).WithAddress(address).WatchTransferEvents(ch)
			}

			block, err := t.database.Blocks().FilterByContractId(contract.Id).Get()
			if err != nil {
				return errors.Wrap(err, "failed to get block to begin with")
			}
			if block != nil {
				startBlock = block.TransferBlock
			}

			listener := t.listener.From(startBlock).WithCtx(ctx).WithAddress(address)
			return listener.WatchTransferEvents(ch)
		},
		t.cfg.Backoff.NormalPeriod,
		t.cfg.Backoff.MinAbnormalPeriod,
		t.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (t *TokenTracker) TrackMintEvents(address common.Address, ch chan<- etherdata.SuccessfulMintEvent) {
	running.WithBackOff(
		t.ctx,
		t.log,
		t.cfg.Prefix+mintTrackerSuffix,
		func(ctx context.Context) error {
			contractEntry, err := t.database.Contracts().GetByAddress(address.String())
			if err != nil {
				return errors.Wrap(err, "failed to get contract from the database")
			}

			if contractEntry == nil {
				t.log.Warnf("The following contract is not contained in the database: %s", address.String())
				return t.listener.From(0).WithCtx(ctx).WithAddress(address).WatchSuccessfulMintEvents(ch)
			}

			return t.listener.From(contractEntry.PreviousMintBLock).WithCtx(ctx).WatchSuccessfulMintEvents(ch)
		},
		t.cfg.Backoff.NormalPeriod,
		t.cfg.Backoff.MinAbnormalPeriod,
		t.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (t *TokenTracker) TrackUpdateEvents(address common.Address, ch chan<- etherdata.UpdateEvent) {
	running.WithBackOff(
		t.ctx,
		t.log,
		t.cfg.Prefix+updateTrackerSuffix,
		func(ctx context.Context) error {
			startBlock := uint64(0)

			contract, err := t.database.Contracts().GetByAddress(address.String())
			if err != nil {
				return errors.Wrap(err, "failed to get contract by address")
			}
			if contract == nil {
				t.log.Warnf("The following contract is not contained in the database: %s", address.String())
				return t.listener.From(0).WithCtx(ctx).WithAddress(address).WatchUpdateEvents(ch)
			}

			block, err := t.database.Blocks().FilterByContractId(contract.Id).Get()
			if err != nil {
				return errors.Wrap(err, "failed to get block to begin with")
			}
			if block != nil {
				startBlock = block.UpdateBlock
			}

			listener := t.listener.From(startBlock).WithCtx(ctx).WithAddress(address)
			return listener.WatchUpdateEvents(ch)
		},
		t.cfg.Backoff.NormalPeriod,
		t.cfg.Backoff.MinAbnormalPeriod,
		t.cfg.Backoff.MaxAbnormalPeriod,
	)
}
