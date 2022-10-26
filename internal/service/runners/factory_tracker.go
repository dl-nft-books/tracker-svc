package runners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/eth_reader"
	"strconv"
)

const factoryTrackerCursor = "factory_tracker_last_block"

type FactoryTracker struct {
	log      *logan.Entry
	database data.DB
	rpc      *ethclient.Client
	reader   eth_reader.FactoryContractReader

	name          string
	address       common.Address
	firstBlock    uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewFactoryTracker(cfg config.Config) *FactoryTracker {
	return &FactoryTracker{
		log:      cfg.Log(),
		database: postgres.NewDB(cfg.DB()),
		rpc:      cfg.EtherClient().Rpc,
		reader:   eth_reader.NewFactoryContractReader(cfg.EtherClient().Rpc),

		name:          cfg.FactoryTracker().Name,
		address:       cfg.FactoryTracker().Address,
		firstBlock:    cfg.FactoryTracker().FirstBlock,
		iterationSize: cfg.FactoryTracker().IterationSize,
		runnerCfg:     cfg.FactoryTracker().Runner,
	}
}

func (t *FactoryTracker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		t.log,
		t.name,
		t.Track,
		t.runnerCfg.NormalPeriod,
		t.runnerCfg.MinAbnormalPeriod,
		t.runnerCfg.MaxAbnormalPeriod,
	)
}

func (t *FactoryTracker) Track(ctx context.Context) error {
	previousBlock, err := t.GetPreviousBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get previous block")
	}

	t.log.Debugf("Trying to iterate from block %d to %d...", previousBlock, previousBlock+t.iterationSize)

	events, _, err := t.reader.GetEvents(t.address, previousBlock, previousBlock+t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}
	if len(events) == 0 {
		t.log.Debug("No create contract events found")
		return nil
	}

	for _, event := range events {
		t.log.Debugf("Caught new contract with block number %d", event.BlockNumber)

		if err = t.InsertEvent(event); err != nil {
			return errors.Wrap(err, "failed to insert event into the database")
		}

		t.log.Debugf("Successfully inserted contract into the database")
	}

	newBlock, err := t.GetNewBlock(previousBlock, t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get new block")
	}

	return t.database.KeyValue().Upsert(data.KeyValue{
		Key:   factoryTrackerCursor,
		Value: strconv.FormatInt(newBlock, 10),
	})
}

func (t *FactoryTracker) GetPreviousBlock() (uint64, error) {
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
	if cursorUInt64 > t.firstBlock {
		return cursorUInt64, nil
	}

	return t.firstBlock, nil
}

func (t *FactoryTracker) GetNewBlock(previousBlock, iterationSize uint64) (int64, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	if previousBlock+iterationSize+1 > lastBlockchainBlock {
		return int64(lastBlockchainBlock) + 1, nil
	}

	return int64(previousBlock + iterationSize + 1), nil
}

func (t *FactoryTracker) InsertEvent(event eth_reader.ContractCreatedEvent) error {
	_, err := t.database.Contracts().Insert(data.Contract{
		Contract:  event.Address.String(),
		Name:      event.Name,
		Symbol:    event.Symbol,
		LastBlock: 0,
	})

	return errors.Wrap(err, "failed to insert contract into the database")
}
