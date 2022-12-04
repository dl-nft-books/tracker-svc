package trackers

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/key_value"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/iterators"
	listeners "gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/listeners"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/processors"
)

type DeployTracker struct {
	log      *logan.Entry
	rpc      *ethclient.Client
	cfg      config.FactoryTracker
	database data.DB

	iterator ethereum.FactoryIterator
	listener ethereum.FactoryListener
}

func NewFactoryTracker(cfg config.Config, ctx context.Context) *DeployTracker {
	var (
		processor = processors.NewFactoryProcessor(cfg, ctx)

		iterator = iterators.NewFactoryIterator(cfg, ctx).
				WithAddress(cfg.FactoryTracker().Address).
				WithProcessor(processor)
		listener = listeners.NewFactoryListener(cfg, ctx).
				WithAddress(cfg.FactoryTracker().Address).
				WithProcessor(processor)
	)

	return &DeployTracker{
		log:      cfg.Log(),
		rpc:      cfg.EtherClient().Rpc,
		cfg:      cfg.FactoryTracker(),
		database: postgres.NewDB(cfg.DB()),

		iterator: iterator,
		listener: listener,
	}
}

func (t *DeployTracker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		t.log,
		t.cfg.Name,
		t.Track,
		t.cfg.Runner.NormalPeriod,
		t.cfg.Runner.MinAbnormalPeriod,
		t.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (t *DeployTracker) Track(ctx context.Context) (err error) {
	startBlock, err := t.GetStartBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get previous block")
	}

	if err = t.iterator.From(startBlock).WithCtx(ctx).ProcessDeployEvents(); err != nil {
		return errors.Wrap(err, "failed to check for events before running the listener")
	}

	if err = t.listener.WithCtx(ctx).WatchContractCreatedEvents(); err != nil {
		return errors.Wrap(err, "failed to watch contract created events")
	}

	return nil
}

// GetStartBlock gets the block to begin with
func (t *DeployTracker) GetStartBlock() (uint64, error) {
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
	if cursorUInt64 > t.cfg.FirstBlock {
		return cursorUInt64, nil
	}

	return t.cfg.FirstBlock, nil
}
