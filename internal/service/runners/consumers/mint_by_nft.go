package consumers

import (
	"context"
	bookerModels "github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/core-svc/connector/models"
	coreResources "github.com/dl-nft-books/core-svc/resources"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"math/big"
	"strconv"
	"time"
)

const nftSymbolKey = "NFT"

func (c *MarketPlaceConsumer) ConsumeMintByNftEvents(ch <-chan etherdata.SuccessfullyMintedByNftEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+mintByNftConsumerSuffix,
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
					check, err := c.database.NftPayments().New().FilterByBookUrl(c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash).Get()
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

					if err = c.MintByNftUpdating(*task, event); err != nil {
						return errors.Wrap(err, "failed to consume mint transaction", logField)
					}

					if err = c.MintByNftStatistics(*book, event); err != nil {
						return errors.Wrap(err, "failed to consume mint transaction", logField)
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

func (c *MarketPlaceConsumer) MintByNftUpdating(task coreResources.Task, event etherdata.SuccessfullyMintedByNftEvent) error {
	// Inserting information about payment
	_, err := c.database.NftPayments().New().Insert(data.NftPayment{
		PayerAddress:      event.Recipient.String(),
		NftAddress:        event.NftAddress.String(),
		NftId:             event.NftId,
		FloorPrice:        event.NftFloorPrice.String(),
		PriceMinted:       event.MintedTokenPrice.String(),
		BookUrl:           c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash,
		ChainId:           c.network.ChainId,
		PurchaseTimestamp: event.Timestamp,
	})

	if err != nil {
		return errors.Wrap(err, "failed to add payment to the table")
	}

	nftRequestList, err := c.core.ListNFTRequest(models.ListNftRequestRequest{
		CollectionAddress: []string{event.NftAddress.String()},
		NftId:             []int64{event.NftId},
		ChainId:           []int64{c.network.ChainId},
		BookId:            []int64{task.Attributes.BookId},
		Status:            []coreResources.NftRequestStatus{coreResources.RequestAccepted},
	})
	// Updating contract`s last mint block
	if err != nil {
		return errors.Wrap(err, "failed to get nft request")
	}
	if nftRequestList != nil {
		if err = c.core.UpdateNFTRequestStatus(cast.ToInt64(nftRequestList.Data[0].ID), coreResources.RequestExchanged); err != nil {
			return errors.Wrap(err, "failed to update nft request status")
		}
	}
	// Updating contract`s last mint block
	if err = c.database.Blocks().UpdateMintBlock(event.BlockNumber, c.network.ChainId); err != nil {
		return errors.Wrap(err, "failed to update contract`s last mint block")
	}

	return nil
}

func (c *MarketPlaceConsumer) MintByNftStatistics(book bookerModels.GetBookResponse, event etherdata.SuccessfullyMintedByNftEvent) error {
	var usdCurrency big.Int
	usdCurrency.Mul(event.MintedTokenPrice, event.NftFloorPrice)
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
			Key:   "stats-book-" + book.Data.ID + "-token_symbol-" + nftSymbolKey + "-price_token",
			Value: (*event.MintedTokenPrice).String(),
		},
		data.KeyValue{ // price (native currency) by each token
			Key:   "stats-token_symbol-" + nftSymbolKey + "-price_token",
			Value: (*event.MintedTokenPrice).String(),
		},
		// USD
		data.KeyValue{ // price (USD) by each book by each token
			Key:   "stats-book-" + book.Data.ID + "-token_symbol-" + nftSymbolKey + "-price_usd",
			Value: usdCurrency.String(),
		}, data.KeyValue{ // price (USD) by each token
			Key:   "stats-token_symbol-" + nftSymbolKey + "-price_usd",
			Value: usdCurrency.String(),
		},
		data.KeyValue{ // price (USD) total
			Key:   "stats-price_usd",
			Value: usdCurrency.String(),
		}, data.KeyValue{ // price (USD) total by each book
			Key:   "stats-book-" + book.Data.ID + "-price_usd",
			Value: usdCurrency.String(),
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
