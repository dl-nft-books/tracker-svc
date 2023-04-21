package consumers

import (
	"context"
	"fmt"
	bookerModels "github.com/dl-nft-books/book-svc/connector/models"
	coreModels "github.com/dl-nft-books/core-svc/connector/models"
	coreResources "github.com/dl-nft-books/core-svc/resources"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/internal/data/opensea"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"math"
	"math/big"
	"strconv"
	"time"
)

func (c *MarketPlaceConsumer) ConsumeTokenSuccessfullyPurchasedEvent(ch <-chan etherdata.TokenSuccessfullyPurchasedEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+mintConsumerSuffix,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": c.network.FactoryAddress}

					// Getting task by hash (uri)
					task, err := c.GetTask(event.Uri)
					if err != nil {
						return errors.Wrap(err, "failed get task")
					}
					if task == nil {
						continue
					}
					logField = logField.Merge(logan.F{
						"task_id": task.ID,
					})

					//Check if Payment with such book_url is already exists
					check, err := c.database.Payments().New().
						FilterByContractAddress(event.ContractAddress.String()).
						FilterByTokenId(task.Attributes.TokenId).
						FilterByChainId(task.Attributes.ChainId).
						Get()
					if err != nil {
						return errors.Wrap(err, "failed to check is payment exist")
					}
					if check != nil {
						c.logger.WithFields(logan.F{"book_url": c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash}).Warn("payment with such book_url is already exist")
						return errors.New("payment with such book_url is already exist")
					}

					book, err := c.GetBook(*task)
					if err != nil {
						return errors.Wrap(err, "failed get book", logField)
					}
					if book == nil {
						continue
					}
					logField = logField.Merge(logan.F{
						"book_id": book.Data.ID,
					})

					if err = c.UploadToIpfs(*book, *task); err != nil {
						return errors.Wrap(err, "failed to upload to IPFS", logField)
					}

					if err = c.MintUpdating(*task, event); err != nil {
						return errors.Wrap(err, "failed to consume mint transaction", logField)
					}

					if err = c.UpdateStatistics(*book, event); err != nil {
						return errors.Wrap(err, "failed to consume mint transaction", logField)
					}

					// Updating contract`s last mint block
					if err = c.database.Blocks().UpdateTokenPurchasedBlockColumn(event.BlockNumber, c.network.ChainId); err != nil {
						return errors.Wrap(err, "failed to update contract`s last mint block")
					}

					c.logger.WithFields(logField).Infof("Successfully processed mint event of a marketplace with id %d", event.TokenId)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (c *MarketPlaceConsumer) GetTask(uri string) (*coreResources.Task, error) {
	// Getting task by hash (uri)
	tasksResponse, err := c.core.ListTasks(coreModels.ListTasksRequest{IpfsHash: &uri})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get task by ipfs hash")
	}
	if len(tasksResponse.Data) == 0 {
		c.logger.Warn("could not find task")
		return nil, nil
	}
	task := tasksResponse.Data[0]

	// Updating status to loading on IPFS
	status := coreResources.TaskUploading
	if err = c.core.UpdateTask(coreModels.UpdateTaskParams{
		Id:     cast.ToInt64(task.ID),
		Status: &status,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update status")
	}
	return &task, nil
}

func (c *MarketPlaceConsumer) GetBook(task coreResources.Task) (*bookerModels.GetBookResponse, error) {
	// Getting book info by task id
	book, err := c.booker.GetBookById(task.Attributes.BookId, task.Attributes.ChainId)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get book of specified task")
	}
	if book == nil {
		c.logger.Warn("could not find book")
		return nil, nil
	}
	return book, nil
}

func (c *MarketPlaceConsumer) UploadToIpfs(book bookerModels.GetBookResponse, task coreResources.Task) error {

	// Getting nft pdf file link
	fileLink, err := c.documenter.GetDocumentLink(book.Data.Attributes.File.Attributes.Key)
	if err != nil {
		return errors.Wrap(err, "failed to get banner image link")
	}

	openseaData := opensea.Metadata{
		Name:        fmt.Sprintf("%s #%d", task.Attributes.TokenName, task.Attributes.TokenId),
		Description: book.Data.Attributes.Description,
		Image:       c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash,
		FileURL:     fileLink.Data.Attributes.Url,
	}
	// Uploading metadata
	if err = c.ipfsLoader.UploadMetadata(openseaData); err != nil {
		return errors.Wrap(err, "failed to load metadata to the ipfs")
	}

	// Uploading banner
	if err = c.ipfsLoader.UploadBanner(task.Attributes.BannerIpfsHash); err != nil {
		return errors.Wrap(err, "failed to load banner to the ipfs")
	}

	taskStatus := coreResources.TaskFinishedUploading
	if err := c.core.UpdateTask(coreModels.UpdateTaskParams{
		Id:     cast.ToInt64(task.ID),
		Status: &taskStatus,
	}); err != nil {
		return errors.Wrap(err, "failed to update task`s marketplace id and status")
	}
	return nil
}

func (c *MarketPlaceConsumer) MintUpdating(task coreResources.Task, event etherdata.TokenSuccessfullyPurchasedEvent) error {
	if _, err := c.database.Payments().New().Insert(data.Payment{
		ContractAddress:   event.ContractAddress.String(),
		TokenId:           task.Attributes.TokenId,
		BookId:            task.Attributes.BookId,
		PayerAddress:      event.Recipient.String(),
		NftId:             event.NftId,
		TokenAddress:      event.Erc20Info.TokenAddress.String(),
		TokenSymbol:       event.Erc20Info.Symbol,
		TokenName:         event.Erc20Info.Name,
		TokenDecimals:     event.Erc20Info.Decimals,
		Amount:            event.Amount.String(),
		PriceMinted:       event.MintedTokenPrice.String(),
		PriceToken:        event.PaymentTokenPrice.String(),
		BannerLink:        c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash,
		PurchaseTimestamp: event.Timestamp,
		ChainId:           c.network.ChainId,
		Type:              int8(event.Type),
	}); err != nil {
		return errors.Wrap(err, "failed to add payment to the table")
	}

	return nil
}

func (c *MarketPlaceConsumer) UpdateStatistics(book bookerModels.GetBookResponse, event etherdata.TokenSuccessfullyPurchasedEvent) error {
	usdPrice := new(big.Float).Quo(new(big.Float).SetInt(event.MintedTokenPrice), big.NewFloat(math.Pow10(18)))
	paymentTokenPrice := new(big.Float).Quo(new(big.Float).SetInt(event.PaymentTokenPrice), big.NewFloat(math.Pow10(18)))
	tokenPrice := big.NewFloat(0).Quo(usdPrice, paymentTokenPrice)
	return c.database.KeyValue().UpdateStatistics(
		// Amount
		data.KeyValue{ // amount of each book
			Key:   "stats-book-" + book.Data.ID + "-amount",
			Value: "1",
		}, data.KeyValue{ // amount of all books
			Key:   "stats-book-amount-total",
			Value: "1",
		}, data.KeyValue{ // amount of all books by day
			Key:   "stats-book-" + book.Data.ID + "-amount-date-" + time.Now().Format("2006-Jan-02"),
			Value: "1",
		},

		// Price
		// Native Currency
		data.KeyValue{ // price (native currency) by each book by each token
			Key:   "stats-book-" + book.Data.ID + "-token_symbol-" + event.Erc20Info.Symbol + "-price_token",
			Value: tokenPrice.String(),
		},
		data.KeyValue{ // price (native currency) by each token
			Key:   "stats-token_symbol-" + event.Erc20Info.Symbol + "-price_token",
			Value: tokenPrice.String(),
		},
		// USD
		data.KeyValue{ // price (USD) by each book by each token
			Key:   "stats-book-" + book.Data.ID + "-token_symbol-" + event.Erc20Info.Symbol + "-price_usd",
			Value: usdPrice.String(),
		}, data.KeyValue{ // price (USD) by each token
			Key:   "stats-token_symbol-" + event.Erc20Info.Symbol + "-price_usd",
			Value: usdPrice.String(),
		},
		data.KeyValue{ // price (USD) total
			Key:   "stats-price_usd",
			Value: usdPrice.String(),
		}, data.KeyValue{ // price (USD) total by each book
			Key:   "stats-book-" + book.Data.ID + "-price_usd",
			Value: usdPrice.String(),
		},

		// Networks
		data.KeyValue{ // network statistic
			Key:   "stats-chain_id-" + strconv.FormatInt(c.network.ChainId, 10),
			Value: "1",
		}, data.KeyValue{ // network statistic by each book
			Key:   "stats-book-" + book.Data.ID + "-chain_id-" + strconv.FormatInt(c.network.ChainId, 10),
			Value: "1",
		})
}
