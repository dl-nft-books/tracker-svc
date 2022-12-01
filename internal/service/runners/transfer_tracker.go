package runners

import (
	"context"
	"strconv"

	"gitlab.com/tokend/nft-books/contract-tracker/internal/contract-reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/contract-reader/evm-based-reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
)

// transferKVTrackerPage is a key that corresponds to the value
// of a page to select contracts according to
const transferKVTrackerPage = "transfer_tracker_page"

var TokenNotFoundErr = errors.New("specified token was not found")

type TransferTracker struct {
	log    *logan.Entry
	rpc    *ethclient.Client
	reader contract_reader.TokenReader
	cfg    config.ContractTracker

	trackerDB   data.TrackerDB
	generatorDB external.GeneratorDB
}

func NewTransferTracker(cfg config.Config) *TransferTracker {
	return &TransferTracker{
		log:    cfg.Log(),
		rpc:    cfg.EtherClient().Rpc,
		reader: evm_based_reader.NewTokenContractReader(cfg),
		cfg:    cfg.TransferTracker(),

		trackerDB:   postgres.NewTrackerDB(cfg.TrackerDB().DB),
		generatorDB: postgres.NewGeneratorDB(cfg.GeneratorDB().DB),
	}
}

func (t *TransferTracker) Run(ctx context.Context) {
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

func (t *TransferTracker) Track(ctx context.Context) error {
	contracts, err := t.FormList()
	if err != nil {
		return errors.Wrap(err, "failed to form a list of contracts")
	}

	for _, contract := range contracts {
		if err = t.ProcessContract(contract, ctx); err != nil {
			return errors.Wrap(err, "failed to process specified contract", logan.F{
				"contract_id": contract.Id,
			})
		}
	}

	return nil
}

func (t *TransferTracker) ProcessContract(contract data.Contract, ctx context.Context) error {
	t.log.Debugf("Processing contract with id of %d...", contract.Id)

	return t.trackerDB.Transaction(func() error {
		lastBlock, err := t.rpc.BlockNumber(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to get last block number")
		}

		startBlock, err := t.GetStartBlock(contract)
		if err != nil {
			return errors.Wrap(err, "failed to get previous block number")
		}

		// Ensuring contract previous block does not exceed last block in the blockchain
		if startBlock >= lastBlock {
			t.log.Debug("Starting block exceeded last block, omitting")
			return nil
		}

		transferEvents, err := t.reader.
			From(startBlock).
			To(startBlock + t.cfg.IterationSize).
			WithAddress(contract.Address()).
			WithCtx(ctx).
			GetTransferEvents()
		if err != nil {
			return errors.Wrap(err, "failed to get successful transfer events")
		}

		if len(transferEvents) == 0 {
			t.log.Debug("No transfer events found")
		}

		for _, event := range transferEvents {
			t.log.Infof("Found transfer events from %s to %s", event.From.String(), event.To.String())

			if err = t.ProcessTransferEvent(event); err != nil {
				return errors.Wrap(err, "failed to process transfer event")
			}

			t.log.Info("Processed transfer event")
		}

		nextBlock := t.GetNextBlock(startBlock, t.cfg.IterationSize, lastBlock)

		if err = t.trackerDB.Blocks().Upsert(data.Blocks{
			ContractId:    contract.Id,
			TransferBlock: nextBlock,
		}); err != nil {
			return errors.Wrap(err, "failed to upsert transfer block", logan.F{
				"transfer_block": nextBlock,
			})
		}

		return nil
	})
}

func (t *TransferTracker) ProcessTransferEvent(event ethereum.TransferEvent) error {
	if event.From == ethereum.NullAddress || event.To == ethereum.NullAddress {
		t.log.Info("Received transfer event with one address being null, omitting")
		return nil
	}

	// FIXME: Update owner via connector with generator service
	token, err := t.generatorDB.Tokens().FilterByTokenId(int64(event.TokenId)).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get token from the database")
	}
	if token == nil {
		t.log.WithFields(logan.F{
			"token_id": event.TokenId,
		}).Warn("token from the event was not found")
		return nil
	}

	if err = t.generatorDB.Tokens().UpdateAccount(event.To.String(), token.Id); err != nil {
		return errors.Wrap(err, "failed to update token's owner", logan.F{
			"new_owner": event.To.String(),
			"token_id":  token.Id,
		})
	}

	return nil
}

func (t *TransferTracker) GetStartBlock(contract data.Contract) (uint64, error) {
	startBlock := uint64(0)

	block, err := t.trackerDB.Blocks().FilterByContractId(contract.Id).Get()
	if err != nil {
		return 0, errors.Wrap(err, "failed to filter by contract id")
	}
	if block != nil {
		startBlock = block.TransferBlock
	}

	return startBlock, nil
}

func (t *TransferTracker) GetNextBlock(startBlock, iterationSize, lastBlock uint64) uint64 {
	if startBlock+iterationSize+1 > lastBlock {
		return lastBlock + 1
	}

	return startBlock + iterationSize + 1
}

func (t *TransferTracker) Select(pageNumber uint64) ([]data.Contract, error) {
	cutQ := t.trackerDB.Contracts().Page(pgdb.OffsetPageParams{
		Limit:      t.cfg.Capacity,
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
