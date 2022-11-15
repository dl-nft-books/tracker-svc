package runners

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/eth_reader"
)

const (
	transferKVTrackerPage = "transfer_tracker_page"
)

var TokenNotFoundErr = errors.New("specified token was not found")

type TransferTracker struct {
	log    *logan.Entry
	rpc    *ethclient.Client
	reader eth_reader.TokenContractReader

	trackerDB   data.TrackerDB
	generatorDB data.GeneratorDB

	name          string
	capacity      uint64
	iterationSize uint64
	runnerCfg     config.Runner
}

func NewTransferTracker(cfg config.Config) *TransferTracker {
	return &TransferTracker{
		log:    cfg.Log(),
		rpc:    cfg.EtherClient().Rpc,
		reader: eth_reader.NewTokenContractReader(cfg.EtherClient().Rpc),

		trackerDB:   postgres.NewTrackerDB(cfg.TrackerDB().DB),
		generatorDB: postgres.NewGeneratorDB(cfg.GeneratorDB().DB),

		name:          cfg.FactoryTracker().Name,
		iterationSize: cfg.FactoryTracker().IterationSize,
		runnerCfg:     cfg.FactoryTracker().Runner,
	}
}

func (t *TransferTracker) Run(ctx context.Context) {
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

func (t *TransferTracker) Track(ctx context.Context) error {
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

func (t *TransferTracker) ProcessContract(contract data.Contract) error {
	t.log.Debugf("Processing contract with id of %d", contract.Id)

	return t.trackerDB.Transaction(func() error {
		previousBlock, err := t.GetPreviousBlock(contract)
		if err != nil {
			return errors.Wrap(err, "failed to get previous block number")
		}

		lastBlock, err := t.rpc.BlockNumber(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to get last block number")
		}

		// Ensuring contract previous block does not exceed last block in the blockchain
		if previousBlock > lastBlock {
			return nil
		}
		// Ensuring previous block does not exceed last mint block
		if previousBlock > contract.LastBlock {
			return nil
		}

		transferEvents, _, err := t.reader.GetTransferEvents(contract.Address(), previousBlock, previousBlock+t.iterationSize)
		if err != nil {
			return errors.Wrap(err, "failed to get successful transfer events")
		}

		if len(transferEvents) == 0 {
			t.log.Debug("No transfer events found")
		}

		for _, event := range transferEvents {
			t.log.Debugf("Found transfer events from %s to %s", event.From.String(), event.To.String())

			if err = t.ProcessTransferEvent(event); err != nil {
				return errors.Wrap(err, "failed to process transfer event")
			}
		}

		newBlock, err := t.GetNewBlock(previousBlock, t.iterationSize)
		if err != nil {
			return errors.Wrap(err, "failed to get new block", logan.F{
				"current_block": contract.LastBlock,
			})
		}

		if err = t.trackerDB.Blocks().Upsert(data.Blocks{
			ContractId:    contract.Id,
			TransferBlock: newBlock,
		}); err != nil {
			return errors.Wrap(err, "failed to upsert transfer block", logan.F{
				"transfer_block": newBlock,
			})
		}

		return nil
	})
}

func (t *TransferTracker) ProcessTransferEvent(event ethereum.TransferEvent) error {
	token, err := t.generatorDB.Tokens().FilterByTokenId(int64(event.TokenId)).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get token from the database")
	}
	if token == nil {
		return errors.From(TokenNotFoundErr, logan.F{
			"token_id": event.TokenId,
		})
	}

	if err = t.generatorDB.Tokens().UpdateAccount(event.To.String(), token.Id); err != nil {
		return errors.Wrap(err, "failed to update token's owner", logan.F{
			"new_owner": event.To.String(),
			"token_id":  token.Id,
		})
	}

	return nil
}

func (t *TransferTracker) GetPreviousBlock(contract data.Contract) (uint64, error) {
	previousBlock := uint64(0)

	block, err := t.trackerDB.Blocks().FilterByContractId(contract.Id).Get()
	if err != nil {
		return 0, errors.Wrap(err, "failed to filter by contract id")
	}
	if block != nil {
		previousBlock = block.TransferBlock
	}

	return previousBlock, nil
}

func (t *TransferTracker) GetNewBlock(previousBlock, iterationSize uint64) (uint64, error) {
	// Retrieving the last blockchain block number
	lastBlockchainBlock, err := t.rpc.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	if previousBlock+iterationSize+1 > lastBlockchainBlock {
		return lastBlockchainBlock + 1, nil
	}

	return previousBlock + iterationSize + 1, nil
}

func (t *TransferTracker) Select(pageNumber uint64) ([]data.Contract, error) {
	cutQ := t.trackerDB.Contracts().Page(pgdb.OffsetPageParams{
		Limit:      t.capacity,
		PageNumber: pageNumber})

	return cutQ.Select()
}

func (t *TransferTracker) FormList() ([]data.Contract, error) {
	pageNumberKV, err := t.trackerDB.KeyValue().Get(transferKVTrackerPage)
	if pageNumberKV == nil {
		pageNumberKV = &data.KeyValue{
			Key:   transferKVTrackerPage,
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
		t.log.Warn("Contracts table is empty")
		return nil, nil
	}

	if len(contracts) == 0 {
		if err = t.trackerDB.KeyValue().Upsert(data.KeyValue{
			Key:   transferKVTrackerPage,
			Value: "0",
		}); err != nil {
			return nil, errors.Wrap(err, "failed to update last processed contract")
		}

		return t.FormList()
	}

	if err = t.trackerDB.KeyValue().Upsert(data.KeyValue{
		Key:   transferKVTrackerPage,
		Value: strconv.FormatInt(pageNumber+1, 10),
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update last processed contract")
	}

	return contracts, nil
}
