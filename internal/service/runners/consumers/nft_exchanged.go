package consumers

import (
	"context"
	bookerModels "github.com/dl-nft-books/book-svc/connector/models"
	coreResources "github.com/dl-nft-books/core-svc/resources"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

func (c *MarketPlaceConsumer) ConsumeTokenSuccessfullyExchangedEvent(ch <-chan etherdata.TokenSuccessfullyExchangedEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+nftExchangedConsumerSuffix,
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

					//Check whether payment with such banner_link already exists
					check, err := c.database.Payments().New().
						FilterByBannerLink(c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash).
						Get()
					if err != nil {
						return errors.Wrap(err, "failed to check is payment exist")
					}
					if check != nil {
						c.logger.WithFields(logan.F{"book_url": c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash}).Warn("payment with such banner_link is already exist")
						continue
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

					if err = c.ExchangeUpdating(*task, event); err != nil {
						return errors.Wrap(err, "failed to consume nft request created transaction", logField)
					}

					if err = c.UpdateStatisticsExchanged(*book, event); err != nil {
						return errors.Wrap(err, "failed to consume nft request created transaction", logField)
					}

					// Updating contract`s last nft request created block
					if err = c.database.Blocks().UpdateTokenExchangedBlockColumn(event.BlockNumber, c.network.ChainId); err != nil {
						return errors.Wrap(err, "failed to update contract`s last nft request created block")
					}

					c.logger.WithFields(logField).Infof("Successfully processed nft request created event of a marketplace with id %d", event.TokenId)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (c *MarketPlaceConsumer) ExchangeUpdating(task coreResources.Task, event etherdata.TokenSuccessfullyExchangedEvent) error {
	if _, err := c.database.Payments().New().Insert(data.Payment{
		ContractAddress:   event.ContractAddress.String(),
		TokenId:           task.Attributes.TokenId,
		BookId:            task.Attributes.BookId,
		PayerAddress:      event.Recipient.String(),
		NftId:             event.NftId,
		TokenAddress:      event.ContractAddress.String(),
		BannerLink:        c.ipfsLoader.BaseUri + task.Attributes.BannerIpfsHash,
		PurchaseTimestamp: event.Timestamp,
		ChainId:           c.network.ChainId,
		Type:              resources.NFT,
	}); err != nil {
		return errors.Wrap(err, "failed to add payment to the table")
	}

	if err := c.database.NftRequests().New().
		UpdateStatus(resources.NftRequestAccepted).
		UpdateLastUpdated(event.Timestamp).
		FilterUpdateById(event.RequestId).Update(); err != nil {
		return errors.Wrap(err, "failed to add payment to the table")
	}

	return nil
}

func (c *MarketPlaceConsumer) UpdateStatisticsExchanged(book bookerModels.GetBookResponse, event etherdata.TokenSuccessfullyExchangedEvent) error {
	bookId := cast.ToInt64(book.Data.ID)

	if err := c.updateBookStatistics(bookId, 0); err != nil {
		return errors.Wrap(err, "failed to update book statistics")
	}
	if err := c.updateBookStatistics(0, 0); err != nil {
		return errors.Wrap(err, "failed to update book statistics")
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
