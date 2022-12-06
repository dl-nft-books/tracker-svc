package trackers

import (
	"context"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/key_value"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/listeners/factory"
)

const deployEventsSuffix = "-factory-deploy"

type FactoryTracker struct {
	log      *logan.Entry
	ctx      context.Context
	database data.DB
	listener ethereum.FactoryListener

	cfg config.Trackers
}

func NewFactoryTracker(cfg config.Config, ctx context.Context) *FactoryTracker {
	return &FactoryTracker{
		log:      cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.Trackers(),
		database: postgres.NewDB(cfg.DB()),
		listener: factory_listeners.NewFactoryListener(cfg, ctx).
			WithAddress(cfg.Trackers().Factory.Address).
			WithMaxDepth(cfg.Trackers().MaxDepth),
	}
}

func (t *FactoryTracker) TrackDeployEvents(ch chan<- etherdata.ContractDeployedEvent) {
	running.WithBackOff(
		t.ctx,
		t.log,
		t.cfg.Prefix+deployEventsSuffix,
		func(ctx context.Context) error {
			startBlock, err := t.GetStartBlock()
			if err != nil {
				return errors.Wrap(err, "failed to get previous block")
			}

			listener := t.listener.From(startBlock).WithCtx(ctx)
			return listener.WatchContractCreatedEvents(ch)
		},
		t.cfg.Backoff.NormalPeriod,
		t.cfg.Backoff.MinAbnormalPeriod,
		t.cfg.Backoff.MaxAbnormalPeriod,
	)
}

// GetStartBlock gets the block to begin with based on config and KV cursor values
func (t *FactoryTracker) GetStartBlock() (uint64, error) {
	cursorKV, err := t.database.KeyValue().Get(key_value.FactoryTrackerCursor)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get cursor value")
	}
	if cursorKV == nil {
		cursorKV = &data.KeyValue{
			Key:   key_value.FactoryTrackerCursor,
			Value: "0",
		}
	}

	cursorUInt64 := cast.ToUint64(cursorKV.Value)
	if cursorUInt64 > t.cfg.Factory.FirstBlock {
		return cursorUInt64, nil
	}

	return t.cfg.Factory.FirstBlock, nil
}
