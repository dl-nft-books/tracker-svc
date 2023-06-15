package consumers

import (
	"context"
	"github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

func (c *MarketPlaceConsumer) ConsumeNFTRequestCreatedEvent(ch <-chan etherdata.NFTRequestCreatedEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+nftRequestCreatedConsumerSuffix,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": c.network.FactoryAddress}

					book, err := c.booker.ListBooks(models.ListBooksParams{
						Contract: []string{event.TokenContract.String()},
						ChainId:  []int64{c.network.ChainId},
					})
					if err != nil {
						c.logger.Error("failed get book")
						return errors.Wrap(err, "failed get book", logField)
					}
					if len(book.Data) == 0 {
						c.logger.Warn("could not find book")
						continue
					}
					if _, err = c.database.NftRequests().New().Insert(data.NftRequest{
						Id:            event.RequestId.Int64(),
						Requester:     event.Requester.String(),
						NftContract:   event.Requester.String(),
						NftId:         event.NftId.Int64(),
						ChainId:       c.network.ChainId,
						BookId:        cast.ToInt64(book.Data[0].ID),
						Status:        resources.NftRequestPending,
						CreatedAt:     event.Timestamp,
						LastUpdatedAt: event.Timestamp,
					}); err != nil {
						return errors.Wrap(err, "failed to consume nft request created transaction", logField)
					}

					// Updating contract`s last nft request created block
					if err = c.database.Blocks().UpdateNFTRequestCreatedBlockColumn(event.BlockNumber, c.network.ChainId); err != nil {
						return errors.Wrap(err, "failed to update contract`s last nft request created block")
					}

					c.logger.WithFields(logField).Infof("Successfully processed nft request created event of a marketplace with id %d", event.RequestId)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}
