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
	"github.com/dl-nft-books/tracker-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"math"
	"math/big"
)

func (c *MarketPlaceConsumer) ConsumeTokenSuccessfullyPurchasedEvent(ch <-chan etherdata.TokenSuccessfullyPurchasedEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+mintConsumerSuffix,
		func(ctx context.Context) (err error) {
			errs := make(chan error, 1)
		mintConsume:
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": c.network.FactoryAddress}

					// Getting task by hash (uri)
					task, err := c.GetTask(event.Uri)
					if err != nil {
						errs <- errors.Wrap(err, "failed get task")
						break mintConsume
					}
					if task == nil {
						continue
					}
					logField = logField.Merge(logan.F{
						"task_id": task.ID,
					})

					//Check whether payment with such banner_link already exists
					check, err := c.database.Payments().New().
						FilterByBannerLink(c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash).
						Get()
					if err != nil {
						errs <- errors.Wrap(err, "failed to check is payment exist")
						break mintConsume
					}
					if check != nil {
						c.logger.WithFields(logan.F{"book_url": c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash}).Warn("payment with such banner_link is already exist")
						continue
					}

					book, err := c.GetBook(*task)
					if err != nil {
						errs <- errors.Wrap(err, "failed get book", logField)
						break mintConsume
					}
					if book == nil {
						continue
					}
					logField = logField.Merge(logan.F{
						"book_id": book.Data.ID,
					})

					if err = c.UploadToIpfs(*book, *task); err != nil {
						errs <- errors.Wrap(err, "failed to upload to IPFS", logField)
						break mintConsume
					}

					if err = c.MintUpdating(*task, event); err != nil {
						errs <- errors.Wrap(err, "failed to consume mint transaction", logField)
						break mintConsume
					}

					if err = c.UpdateStatistics(*book, event); err != nil {
						errs <- errors.Wrap(err, "failed to consume mint transaction", logField)
						break mintConsume
					}

					// Updating contract`s last mint block
					if err = c.database.Blocks().New().UpdateTokenPurchasedBlockColumn(event.BlockNumber, c.network.ChainId); err != nil {
						errs <- errors.Wrap(err, "failed to update contract`s last mint block")
						break mintConsume
					}

					c.logger.WithFields(logField).Infof("Successfully processed mint event of a marketplace with id %d", event.TokenId)
				}
			}

			if err, open := <-errs; open {
				return err
			}
			return nil
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
		Type:              resources.TokenPurchasedEventType(event.Type),
	}); err != nil {
		return errors.Wrap(err, "failed to add payment to the table")
	}

	return nil
}

func (c *MarketPlaceConsumer) UpdateStatistics(book bookerModels.GetBookResponse, event etherdata.TokenSuccessfullyPurchasedEvent) error {
	pricePerOneToken, _ := new(big.Float).Quo(new(big.Float).SetInt(event.PaymentTokenPrice), big.NewFloat(math.Pow10(18))).Float64()
	tokenPrice, _ := new(big.Float).Quo(new(big.Float).SetInt(event.Amount), big.NewFloat(math.Pow10(18))).Float64()
	usdPrice := pricePerOneToken * tokenPrice
	bookId := cast.ToInt64(book.Data.ID)

	if err := c.updateBookStatistics(bookId, usdPrice); err != nil {
		return errors.Wrap(err, "failed to update book statistics")
	}
	if err := c.updateBookStatistics(0, usdPrice); err != nil {
		return errors.Wrap(err, "failed to update book statistics")
	}

	if err := c.updateTokenStatistics(bookId, usdPrice, tokenPrice, event.Erc20Info.Symbol); err != nil {
		return errors.Wrap(err, "failed to update token statistics ")
	}
	if err := c.updateTokenStatistics(0, usdPrice, tokenPrice, event.Erc20Info.Symbol); err != nil {
		return errors.Wrap(err, "failed to update token statistics ")
	}

	if err := c.updateDateStatistics(bookId, event.Timestamp.Format(data.TimestampFormat)); err != nil {
		return errors.Wrap(err, "failed to update date statistics ")
	}

	if err := c.updateChainStatistics(bookId); err != nil {
		return errors.Wrap(err, "failed to update chain statistics ")
	}
	if err := c.updateChainStatistics(0); err != nil {
		return errors.Wrap(err, "failed to update chain statistics ")
	}

	return nil
}

func (c *MarketPlaceConsumer) updateBookStatistics(bookId int64, bookPrice float64) error {
	bookStats, err := c.database.Statistics().BookStatisticsQ.New().FilterByBookId(bookId).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get statistics by book")
	}
	if bookStats != nil {
		return c.database.Statistics().BookStatisticsQ.New().Update(data.BookStatistics{
			Amount:   bookStats.Amount + 1,
			UsdPrice: bookStats.UsdPrice + bookPrice,
		}, bookStats.Id)
	}
	_, err = c.database.Statistics().BookStatisticsQ.New().Insert(data.BookStatistics{
		Amount:   1,
		UsdPrice: bookPrice,
		BookId:   bookId,
	})
	return err
}

func (c *MarketPlaceConsumer) updateTokenStatistics(bookId int64, usdPrice, tokenPrice float64, symbol string) error {
	tokenStats, err := c.database.Statistics().TokenStatisticsQ.New().FilterByBookId(bookId).FilterByTokenSymbol(symbol).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get statistics by token")
	}
	if tokenStats != nil {
		return c.database.Statistics().TokenStatisticsQ.New().Update(data.TokenStatistics{
			UsdPrice:   tokenStats.UsdPrice + usdPrice,
			TokenPrice: tokenStats.TokenPrice + tokenPrice,
		}, tokenStats.Id)
	}
	_, err = c.database.Statistics().TokenStatisticsQ.New().Insert(data.TokenStatistics{
		TokenSymbol: symbol,
		UsdPrice:    usdPrice,
		TokenPrice:  tokenPrice,
		BookId:      bookId,
	})
	return err
}

func (c *MarketPlaceConsumer) updateDateStatistics(bookId int64, date string) error {
	dateStats, err := c.database.Statistics().DateStatisticsQ.New().
		FilterByBookId(bookId).
		FilterByDate(date).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get statistics by date")
	}
	if dateStats != nil {
		return c.database.Statistics().DateStatisticsQ.New().Update(dateStats.Amount+1, dateStats.Id)
	}
	_, err = c.database.Statistics().DateStatisticsQ.New().Insert(data.DateStatistics{
		Amount: 1,
		Date:   date,
		BookId: bookId,
	})
	return err
}

func (c *MarketPlaceConsumer) updateChainStatistics(bookId int64) error {
	chainStats, err := c.database.Statistics().ChainStatisticsQ.New().
		FilterByBookId(bookId).FilterByChainId(c.network.ChainId).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get statistics by chain")
	}
	if chainStats != nil {
		return c.database.Statistics().ChainStatisticsQ.New().Update(chainStats.Amount+1, chainStats.Id)
	}
	_, err = c.database.Statistics().ChainStatisticsQ.New().Insert(data.ChainStatistics{
		Amount:  1,
		ChainId: c.network.ChainId,
		BookId:  bookId,
	})
	return err
}
