package runners

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/reader/ethreader"
	"strconv"
)

const factoryTrackerCursor = "factory_tracker_last_block"

type FactoryTracker struct {
	log      *logan.Entry
	rpc      *ethclient.Client
	cfg      config.FactoryTracker
	database data.TrackerDB
	reader   reader.FactoryReader
}

func NewFactoryTracker(cfg config.Config) *FactoryTracker {
	return &FactoryTracker{
		log:      cfg.Log(),
		rpc:      cfg.EtherClient().Rpc,
		cfg:      cfg.FactoryTracker(),
		database: postgres.NewTrackerDB(cfg.TrackerDB().DB),
		reader:   ethreader.NewFactoryContractReader(cfg).WithAddress(cfg.FactoryTracker().Address),
	}
}

func (t *FactoryTracker) Run(ctx context.Context) {
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

func (t *FactoryTracker) Track(ctx context.Context) error {
	lastBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get last block number")
	}

	startBlock, err := t.GetStartBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get previous block")
	}

	if startBlock >= lastBlock {
		t.log.Debug("Starting block exceeded last block, omitting")
		return nil
	}

	t.log.Debugf("Trying to iterate from block %d to %d...", startBlock, startBlock+t.cfg.IterationSize)

	events, err := t.reader.
		From(startBlock).
		To(startBlock + t.cfg.IterationSize).
		GetContractCreatedEvents()
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	for _, event := range events {
		t.log.Infof("Caught newly deployed contract with block number %d", event.BlockNumber)

		if err = t.InsertEvent(event); err != nil {
			return errors.Wrap(err, "failed to insert event into the database")
		}

		t.log.Debugf("Successfully inserted contract into the database")
	}

	nextBlock, err := t.GetNextBlock(startBlock, t.cfg.IterationSize, lastBlock)
	if err != nil {
		return errors.Wrap(err, "failed to get new block")
	}

	if err = t.database.KeyValue().Upsert(data.KeyValue{
		Key:   factoryTrackerCursor,
		Value: strconv.FormatInt(nextBlock, 10),
	}); err != nil {
		return errors.Wrap(err, "failed to upsert new value")
	}

	return nil
}

// GetStartBlock gets the block to begin with
func (t *FactoryTracker) GetStartBlock() (uint64, error) {
	cursorKV, err := t.database.KeyValue().Get(factoryTrackerCursor)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get cursor value")
	}
	if cursorKV == nil {
		cursorKV = &data.KeyValue{
			Key:   factoryTrackerCursor,
			Value: "0",
		}
	}

	cursor, err := strconv.ParseInt(cursorKV.Value, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert cursor value from string to integer")
	}

	cursorUInt64 := uint64(cursor)
	if cursorUInt64 > t.cfg.FirstBlock {
		return cursorUInt64, nil
	}

	return t.cfg.FirstBlock, nil
}

func (t *FactoryTracker) GetNextBlock(startBlock, iterationSize, lastBlock uint64) (int64, error) {
	if startBlock+iterationSize+1 > lastBlock {
		return int64(lastBlock + 1), nil
	}

	return int64(startBlock + iterationSize + 1), nil
}

func (t *FactoryTracker) InsertEvent(event ethereum.ContractCreatedEvent) error {
	_, err := t.database.Contracts().Insert(data.Contract{
		Contract:  event.Address.String(),
		Name:      event.Name,
		Symbol:    event.Symbol,
		LastBlock: event.BlockNumber,
	})

	return errors.Wrap(err, "failed to insert contract into the database")
}
