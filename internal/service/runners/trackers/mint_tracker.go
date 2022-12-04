package trackers

/*
import (
	"context"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/opensea"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/iterators"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

const (
	mintKVTrackerPage = "mint_tracker_page"
	baseURI           = "https://ipfs.io/ipfs/"
)

var bannerNotFoundErr = errors.New("specified banner key was not found")

type MintTracker struct {
	log        *logan.Entry
	rpc        *ethclient.Client
	reader     ethereum.TokenReader
	ipfsLoader *helpers.IpfsLoader
	cfg        config.ContractTracker

	trackerDB   data.TrackerDB
	generatorDB external.GeneratorDB
	booksQ      external.BookQ

	documenter *documenter.Connector
}

func NewMintTracker(cfg config.Config) *MintTracker {
	return &MintTracker{
		log:        cfg.Log(),
		rpc:        cfg.EtherClient().Rpc,
		reader:     iterators.NewTokenContractReader(cfg),
		ipfsLoader: helpers.NewIpfsLoader(cfg),
		cfg:        cfg.MintTracker(),

		trackerDB:   postgres.NewTrackerDB(cfg.TrackerDB().DB),
		generatorDB: postgres.NewGeneratorDB(cfg.GeneratorDB().DB),
		booksQ:      postgres.NewBooksQ(cfg.BookDB().DB),

		documenter: cfg.DocumenterConnector(),
	}
}

func (t *MintTracker) Run(ctx context.Context) {
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

func (t *MintTracker) Track(ctx context.Context) error {
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

func (t *MintTracker) ProcessContract(contract data.Contract, ctx context.Context) error {
	t.log.Debugf("Processing contract with id of %d...", contract.Id)

	return t.trackerDB.Transaction(func() error {
		// Getting last block in the blockchain
		lastBlock, err := t.rpc.BlockNumber(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to get last block number")
		}

		successfulMintEvents, err := t.reader.
			From(contract.LastBlock).
			To(lastBlock).
			WithAddress(contract.Address()).
			WithCtx(ctx).
			GetSuccessfulMintEvents()
		if err != nil {
			return errors.Wrap(err, "failed to get successful mint events")
		}

		if len(successfulMintEvents) == 0 {
			t.log.Debug("No successful events found")
		}

		for _, event := range successfulMintEvents {
			t.log.Infof("Found successful mint event with a block number of %d", event.BlockNumber)

			if err = t.ProcessSuccessfulMintEvent(contract, event); err != nil {
				return errors.Wrap(err, "failed to process successful mint event")
			}

			t.log.Info("Processed successful mint event")
		}

		return nil
	})
}

func (t *MintTracker) GetNextBlock(startBlock, iterationSize, lastBlock uint64) uint64 {
	if startBlock+iterationSize+1 > lastBlock {
		return lastBlock + 1
	}

	return startBlock + iterationSize + 1
}

func (t *MintTracker) ProcessSuccessfulMintEvent(contract data.Contract, event etherdata.SuccessfulMintEvent) error {
	return t.trackerDB.Transaction(func() error {
		// FIXME: Make the following actions via connectors:
		// 1. Get task info using event uri
		// 2. Get book info using retrieved from step 1 (e.g., via included)
		// 3. Insert new token into generator table
		// 4. Update status of a task

		// Updating book and tasks db
		t.booksQ = t.booksQ.New()

		// Getting task by hash (uri)
		task, err := t.generatorDB.Tasks().GetByHash(event.Uri)
		if err != nil {
			return errors.Wrap(err, "failed to get task by ipfs hash")
		}
		if task == nil {
			t.log.Warnf("Could not find task with a hash of %s", event.Uri)
			return nil
		}

		// Getting book info
		book, err := t.booksQ.FilterActual().FilterByID(task.BookId).Get()
		if err != nil {
			return errors.Wrap(err, "failed to get book by task id")
		}
		if book == nil {
			t.log.WithFields(logan.F{"book_id": task.BookId}).Warn("Could not find book")
			return nil
		}

		// Parsing banner key
		bannerKey, err := helpers.GetDocumentKey(book.Banner)
		if err != nil {
			return errors.Wrap(err, "failed to get document key")
		}
		if bannerKey == nil {
			return bannerNotFoundErr
		}

		// Getting nft banner img link
		bannerLink, err := t.documenter.GetDocumentLink(*bannerKey)
		if err != nil {
			return errors.Wrap(err, "failed to get banner image link")
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
			PriceMinted:       event.MintedTokenPrice.String(),
			PriceToken:        event.PaymentTokenPrice.String(),
			PurchaseTimestamp: event.Timestamp,
			BookUrl:           baseURI + task.FileIpfsHash,
		})
		if err != nil {
			return errors.Wrap(err, "failed to add payment to the table")
		}

		// Inserting information about token
		tokenId, err := t.generatorDB.Tokens().Insert(external.Token{
			Account:      event.Recipient.String(),
			TokenId:      event.TokenId,
			BookId:       book.ID,
			PaymentId:    paymentId,
			MetadataHash: task.MetadataIpfsHash,
			Signature:    task.Signature,
			Status:       resources.TokenUploading,
		})

		// Uploading metadata
		if err = t.ipfsLoader.UploadMetadata(opensea.Metadata{
			Name:        fmt.Sprintf("%s #%v", book.Title, task.Id),
			Description: book.Description,
			Image:       bannerLink.Data.Attributes.Url,
			FileURL:     baseURI + task.FileIpfsHash,
		}); err != nil {
			return errors.Wrap(err, "failed to load metadata to the ipfs")
		}

		// Uploading file
		if err = t.ipfsLoader.UploadFile(task.FileIpfsHash); err != nil {
			return errors.Wrap(err, "failed to load file to the ipfs")
		}

		// Updating tokenID
		if err = t.generatorDB.Tasks().UpdateTokenId(event.TokenId, task.Id); err != nil {
			return errors.Wrap(err, "failed to update token id")
		}

		// Updating status to fully processed task
		if err = t.generatorDB.Tasks().UpdateStatus(resources.TaskFinishedUploading, task.Id); err != nil {
			return errors.Wrap(err, "failed to update task's status", logan.F{
				"task_id": task.Id,
			})
		}

		// Updating status to a token
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
		Limit:      t.cfg.Capacity,
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
*/
