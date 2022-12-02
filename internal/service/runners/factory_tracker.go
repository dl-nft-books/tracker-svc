package runners

import (
	"context"
	"github.com/spf13/cast"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/iterators"
)

const factoryTrackerCursor = "factory_tracker_last_block"

type FactoryTracker struct {
	log      *logan.Entry
	rpc      *ethclient.Client
	cfg      config.FactoryTracker
	database data.DB
	reader   ethereum.FactoryReader
}

func NewFactoryTracker(cfg config.Config) *FactoryTracker {
	return &FactoryTracker{
		log:      cfg.Log(),
		rpc:      cfg.EtherClient().Rpc,
		cfg:      cfg.FactoryTracker(),
		database: postgres.NewDB(cfg.DB()),
		reader:   iterators.NewFactoryContractReader(cfg).WithAddress(cfg.FactoryTracker().Address),
	}
}

// TODO: ADD PROCESS EVENT TO UPDATE BOOK PARAMS
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

func (t *FactoryTracker) Track(ctx context.Context) (err error) {
	startBlock, err := t.GetStartBlock()
	if err != nil {
		return errors.Wrap(err, "failed to get previous block")
	}

	t.log.Debug("Trying to iterate through blocks...")

	if err = t.database.Transaction(func() error {
		events, err := t.reader.From(startBlock).GetContractCreatedEvents()
		if err != nil {
			return errors.Wrap(err, "failed to get events")
		}

		if len(events) == 0 {
			t.log.Debug("No deploy events found")
		}

		contractsBatch := formContractsBatch(events)
		if _, err = t.database.Contracts().Insert(contractsBatch...); err != nil {
			return errors.Wrap(err, "failed to insert contracts batch into the database")
		}

		t.log.Debug("Successfully inserted contract batch into the database")

		// Getting last blockchain block
		lastChainBlock, err := t.rpc.BlockNumber(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get last block number")
		}

		return t.database.KeyValue().Upsert(data.KeyValue{
			Key:   factoryTrackerCursor,
			Value: strconv.FormatUint(lastChainBlock, 10),
		})
	}); err != nil {
		return errors.Wrap(err, "failed to perform a database tx")
	}

	events := make(chan etherdata.ContractDeployedEvent)
	go func() {
		if tempErr := t.reader.WatchContractCreatedEvents(events); tempErr != nil {
			err = tempErr
			return
		}
	}()
	go func() {
		for {
			select {}
		}
	}
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

	cursorUInt64 := cast.ToUint64(cursorKV.Value)
	if cursorUInt64 > t.cfg.FirstBlock {
		return cursorUInt64, nil
	}

	return t.cfg.FirstBlock, nil
}

func formContractsBatch(events []etherdata.ContractDeployedEvent) (batch []data.Contract) {
	for _, event := range events {
		batch = append(batch, data.Contract{
			Contract:  event.Address.String(),
			Name:      event.Name,
			Symbol:    event.Symbol,
			LastBlock: event.BlockNumber,
		})
	}

	return
}

// TODO: ADD PROCESS EVENT TO UPDATE BOOK PARAMS
func (t *FactoryTracker) ProcessEvent(event etherdata.ContractDeployedEvent) error {

	//book, err := t.database.Books().New().FilterByTokenId(int64(event.TokenId)).Get()
	//if err != nil {
	//	return errors.Wrap(err, "failed to get book by token id")
	//}
	//if book == nil {
	//	t.log.Warnf("Book with token id %v was not found", event.TokenId)
	//	return nil
	//}
	//
	//switch event.Status {
	//case types.ReceiptStatusSuccessful:
	//	return t.database.Transaction(
	//		func() error {
	//			if err = t.database.Books().UpdateContractAddress(event.Address.String(), book.ID); err != nil {
	//				return errors.Wrap(err, "failed to update contract address", logan.F{
	//					"contract_address": event.Address.String(),
	//				})
	//			}
	//
	//			return t.database.Books().UpdateDeployStatus(resources.DeploySuccessful, book.ID)
	//		})
	//case types.ReceiptStatusFailed:
	//	return t.database.Books().UpdateDeployStatus(resources.DeployFailed, book.ID)
	//}
	//
	return errors.From(errors.New(""), logan.F{
		"block_number": event.BlockNumber,
	})
}
