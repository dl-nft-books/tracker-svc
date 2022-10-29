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
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
	"strconv"
)

const mintKVTrackerPage = "mint_tracker_page"

type MintTracker struct {
	log        *logan.Entry
	rpc        *ethclient.Client
	reader     eth_reader.TokenContractReader
	ipfsLoader *ipfs_loader.IpfsLoader

	trackerDB data.TrackerDB
	tasksQ    data.TasksQ

	name          string
	capacity      uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewMintTracker(cfg config.Config) *MintTracker {
	return &MintTracker{
		log:        cfg.Log(),
		rpc:        cfg.EtherClient().Rpc,
		reader:     eth_reader.NewTokenContractReader(cfg.EtherClient().Rpc),
		ipfsLoader: ipfs_loader.NewIpfsLoader(cfg),

		trackerDB: postgres.NewTrackerDB(cfg.TrackerDB().DB),
		tasksQ:    postgres.NewTasksQ(cfg.GeneratorDB().DB),

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
	lastBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get last block number")
	}

	if contract.LastBlock > lastBlock {
		t.log.Debugf("contract last block exceeded last block in the blockchain")
		return nil
	}
	events, _, err := t.reader.GetEvents(contract.Address(), contract.LastBlock, contract.LastBlock+t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get events")
	}

	if len(events) == 0 {
		t.log.Debug("No events found")
	}

	for _, event := range events {
		t.log.Debugf("Found event with a block number of %d and uri %s", event.BlockNumber, event.Uri)

		if err = t.ProcessEvent(event); err != nil {
			return errors.Wrap(err, "failed to process event", logan.F{
				"event_block_number": event.BlockNumber,
				"event_uri":          event.Uri,
			})
		}
	}

	newBlock, err := t.GetNewBlock(contract.LastBlock, t.iterationSize)
	if err != nil {
		return errors.Wrap(err, "failed to get new block", logan.F{
			"current_block": contract.LastBlock,
		})
	}

	if err = t.trackerDB.Contracts().UpdateLastBlock(newBlock, contract.Id); err != nil {
		return errors.Wrap(err, "failed to update last block")
	}

	return nil
}

func (t *MintTracker) GetNewBlock(previousBlock, iterationSize uint64) (uint64, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	t.log.Debugf("Last blockchain block has id of %d", lastBlockchainBlock)

	if previousBlock+iterationSize+1 > lastBlockchainBlock {
		return lastBlockchainBlock + 1, nil
	}

	return previousBlock + iterationSize + 1, nil
}

func (t *MintTracker) ProcessEvent(event eth_reader.TokenMintedEvent) error {
	ipfsHash := t.ipfsLoader.GetHashOutUri(event.Uri)

	return t.trackerDB.Transaction(func() error {
		task, err := t.tasksQ.GetByHash(ipfsHash)
		if err != nil {
			return errors.Wrap(err, "failed to get task by ipfs hash")
		}
		if task == nil {
			t.log.Warnf("Could not find task with a hash of %s", ipfsHash)
			return nil
		}

		if err = t.tasksQ.UpdateStatus(resources.TaskUploading, task.Id); err != nil {
			return errors.Wrap(err, "failed to update status", logan.F{
				"task_id": task.Id,
			})
		}

		if err = t.ipfsLoader.Load(event.Uri); err != nil {
			return errors.Wrap(err, "failed to load file to the ipfs")
		}

		if err = t.tasksQ.UpdateIpfsHash(ipfsHash, task.Id); err != nil {
			return errors.Wrap(err, "failed to update ipfs hash")
		}

		if err = t.tasksQ.UpdateTokenId(event.TokenId, task.Id); err != nil {
			return errors.Wrap(err, "failed to update token id")
		}

		if err = t.tasksQ.UpdateStatus(resources.TaskFinishedUploading, task.Id); err != nil {
			return errors.Wrap(err, "failed to update status", logan.F{
				"task_id": task.Id,
			})
		}
		return nil
	})
}

func (t *MintTracker) Select(pageNumber uint64) ([]data.Contract, error) {
	cutQ := t.trackerDB.Contracts().Page(pgdb.OffsetPageParams{
		Limit:      t.capacity,
		PageNumber: pageNumber})

	return cutQ.Select()
}

func (t *MintTracker) FormList() ([]data.Contract, error) {
	pageNumberKV, err := t.trackerDB.KeyValue().Get(mintKVTrackerPage)
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
		if err = t.trackerDB.KeyValue().Upsert(data.KeyValue{
			Key:   mintKVTrackerPage,
			Value: "0",
		}); err != nil {
			return nil, errors.Wrap(err, "failed to update last processed contract")
		}

		return t.FormList()
	}

	if err = t.trackerDB.KeyValue().Upsert(data.KeyValue{
		Key:   mintKVTrackerPage,
		Value: strconv.FormatInt(pageNumber+1, 10),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update last processed contract")
	}

	return contracts, nil
}
