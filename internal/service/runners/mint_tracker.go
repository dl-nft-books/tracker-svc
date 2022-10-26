package runners

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/eth_reader"
	"strconv"
)

const mintKVTrackerPage = "mint_tracker_page"

type MintTracker struct {
	log      *logan.Entry
	database data.DB
	rpc      *ethclient.Client
	reader   eth_reader.TokenContractReader

	name          string
	capacity      uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewMintTracker(cfg config.Config) *MintTracker {
	return &MintTracker{
		log:      cfg.Log(),
		database: postgres.NewDB(cfg.DB()),
		rpc:      cfg.EtherClient().Rpc,
		reader:   eth_reader.NewTokenContractReader(cfg.EtherClient().Rpc),

		name:          cfg.MintTracker().Name,
		iterationSize: cfg.MintTracker().IterationSize,
		runnerCfg:     cfg.MintTracker().Runner,
	}
}

func (t *MintTracker) Run(ctx context.Context) {
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

func (t *MintTracker) Track(ctx context.Context) error {
	contracts, err := t.FormList()
	if err != nil {
		return errors.Wrap(err, "failed to form a list of contracts")
	}

	for _, contract := range contracts {
		if err = t.ProcessContract(contract); err != nil {
			return errors.Wrap(err, "failed to process specified contract", logan.F{
				"contract_id": contract.Id,
			})
		}
	}

	return nil
}

func (t *MintTracker) ProcessContract(contract data.Contract) error {
	t.log.Debugf("Processing contract with id of %d", contract.Id)

	events, _, err := t.reader.GetEvents(contract.Address(), contract.LastBlock, contract.LastBlock+t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	if len(events) == 0 {
		t.log.Debug("No events found")
	}

	for _, event := range events {
		t.log.Debugf("Found event with a block number of %d", event.BlockNumber)
		// TODO: Add event processing here
	}

	return nil
}

func (t *MintTracker) Select(pageNumber uint64) ([]data.Contract, error) {
	cutQ := t.database.Contracts().Page(pgdb.OffsetPageParams{
		Limit:      t.capacity,
		PageNumber: pageNumber})

	return cutQ.Select()
}

func (t *MintTracker) FormList() ([]data.Contract, error) {
	pageNumberKV, err := t.database.KeyValue().Get(mintKVTrackerPage)
	if pageNumberKV == nil {
		pageNumberKV = &data.KeyValue{
			Key:   mintKVTrackerPage,
			Value: "0",
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get page number from KV table")
	}

	pageNumber, err := strconv.ParseInt(pageNumberKV.Value, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert page number to an integer format")
	}

	contracts, err := t.Select(uint64(pageNumber))
	if err != nil {
		return nil, errors.Wrap(err, "failed to select contracts from the table")
	}

	if len(contracts) == 0 && pageNumber == 0 {
		t.log.Warn("contracts table is empty")
		return nil, nil
	}

	if len(contracts) == 0 {
		if err = t.database.KeyValue().Upsert(data.KeyValue{
			Key:   mintKVTrackerPage,
			Value: "0",
		}); err != nil {
			return nil, errors.Wrap(err, "failed to update last processed contract")
		}

		return t.FormList()
	}

	if err = t.database.KeyValue().Upsert(data.KeyValue{
		Key:   mintKVTrackerPage,
		Value: strconv.FormatInt(pageNumber+1, 10),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update last processed contract")
	}

	return contracts, nil
}
