package runners

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/eth_reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/models"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

const (
	mintKVTrackerPage = "mint_tracker_page"
	baseURI           = "https://ipfs.io/ipfs/"
)

type MintTracker struct {
	log        *logan.Entry
	rpc        *ethclient.Client
	reader     eth_reader.TokenContractReader
	ipfsLoader *helpers.IpfsLoader

	trackerDB   data.TrackerDB
	generatorDB data.GeneratorDB
	booksQ      data.BookQ

	name          string
	capacity      uint64
	iterationSize uint64
	runnerCfg     config.Runner

	documenterConnector *s3connector.Connector
}

func NewMintTracker(cfg config.Config) *MintTracker {
	return &MintTracker{
		log:        cfg.Log(),
		rpc:        cfg.EtherClient().Rpc,
		reader:     eth_reader.NewTokenContractReader(cfg.EtherClient().Rpc),
		ipfsLoader: helpers.NewIpfsLoader(cfg),

		trackerDB:   postgres.NewTrackerDB(cfg.TrackerDB().DB),
		generatorDB: postgres.NewGeneratorDB(cfg.GeneratorDB().DB),
		booksQ:      postgres.NewBooksQ(cfg.BookDB().DB),

		name:          cfg.MintTracker().Name,
		iterationSize: cfg.MintTracker().IterationSize,
		runnerCfg:     cfg.MintTracker().Runner,

		documenterConnector: cfg.DocumenterConnector(),
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

	return t.trackerDB.Transaction(func() error {
		lastBlock, err := t.rpc.BlockNumber(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to get last block number")
		}

		if contract.LastBlock > lastBlock {
			return nil
		}

		successfulMintEvents, _, err := t.reader.GetSuccessfulMintEvents(contract.Address(), contract.LastBlock, contract.LastBlock+t.iterationSize)
		if err != nil {
			return errors.Wrap(err, "failed to get successful mint events")
		}

		if len(successfulMintEvents) == 0 {
			t.log.Debug("No successful events found")
		}

		for _, event := range successfulMintEvents {
			t.log.Debugf("Found successful mint event with a block number of %d", event.BlockNumber)

			if err = t.ProcessSuccessfulMintEvent(contract, event); err != nil {
				return errors.Wrap(err, "failed to process successful mint event")
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
	})
}

func (t *MintTracker) GetNewBlock(previousBlock, iterationSize uint64) (uint64, error) {
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

func (t *MintTracker) ProcessSuccessfulMintEvent(contract data.Contract, event eth_reader.SuccessfulMintEvent) error {
	return t.trackerDB.Transaction(func() error {
		// updating book and tasks db
		t.booksQ = t.booksQ.New()

		// event.Uri is actually metadata hash
		task, err := t.generatorDB.Tasks().GetByHash(event.Uri)
		if err != nil {
			return errors.Wrap(err, "failed to get task by ipfs hash")
		}
		if task == nil {
			t.log.Warnf("Could not find task with a hash of %s", event.Uri)
			return nil
		}

		// getting book
		book, err := t.booksQ.FilterActual().FilterByID(task.BookId).Get()
		if err != nil {
			return errors.Wrap(err, "failed to get book by task id")
		}
		if book == nil {
			t.log.Warnf("Could not find book with id %d", task.BookId)
			return nil
		}

		// parsing banner key
		bannerKey, err := helpers.GetDocumentKey(book.Banner)
		if err != nil {
			return err
		}
		if bannerKey == nil {
			return errors.Wrap(err, "failed to get document key")
		}

		// getting nft banner img link
		bannerLink, err := t.documenterConnector.GetDocumentLink(*bannerKey)
		if err != nil {
			return err
		}

		// updating status to loading on IPFS
		if err = t.generatorDB.Tasks().UpdateStatus(resources.TaskUploading, task.Id); err != nil {
			return errors.Wrap(err, "failed to update status", logan.F{
				"task_id": task.Id,
			})
		}

		// inserting information about payment
		paymentId, err := t.trackerDB.Payments().Insert(data.Payment{
			ContractId:        contract.Id,
			ContractAddress:   contract.Contract,
			PayerAddress:      event.Recipient.String(),
			TokenAddress:      event.Erc20Info.TokenAddress.String(),
			TokenSymbol:       event.Erc20Info.Symbol,
			TokenName:         event.Erc20Info.Name,
			TokenDecimals:     event.Erc20Info.Decimals,
			Amount:            event.Amount.String(),
			Price:             event.Price.String(),
			PurchaseTimestamp: event.Timestamp,
			BookUrl:           baseURI + task.FileIpfsHash,
		})
		if err != nil {
			return errors.Wrap(err, "failed to add payment to the table")
		}

		// inserting information about token
		tokenId, err := t.generatorDB.Tokens().Insert(data.Token{
			Account:      event.Recipient.String(),
			TokenId:      event.TokenId,
			BookId:       book.ID,
			PaymentId:    paymentId,
			MetadataHash: task.MetadataIpfsHash,
			Signature:    task.Signature,
			Status:       resources.TokenUploading,
		})

		// uploading metadata
		if err = t.ipfsLoader.UploadMetadata(models.Metadata{
			Name:        fmt.Sprintf("%s #%v", book.Title, task.Id),
			Description: book.Description,
			Image:       bannerLink.Data.Attributes.Url,
			FileURL:     baseURI + task.FileIpfsHash,
		}); err != nil {
			return errors.Wrap(err, "failed to load metadata to the ipfs")
		}

		// uploading file
		if err = t.ipfsLoader.UploadFile(task.FileIpfsHash); err != nil {
			return errors.Wrap(err, "failed to load file to the ipfs")
		}

		// updating tokenID
		if err = t.generatorDB.Tasks().UpdateTokenId(event.TokenId, task.Id); err != nil {
			return errors.Wrap(err, "failed to update token id")
		}

		// updating status to fully processed task
		if err = t.generatorDB.Tasks().UpdateStatus(resources.TaskFinishedUploading, task.Id); err != nil {
			return errors.Wrap(err, "failed to update task's status", logan.F{
				"task_id": task.Id,
			})
		}

		// updating status to a token
		if err = t.generatorDB.Tokens().UpdateStatus(resources.TokenFinishedUploading, tokenId); err != nil {
			return errors.Wrap(err, "failed to update token's status", logan.F{
				"token_id": tokenId,
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
		t.log.Warn("Contracts table is empty")
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
