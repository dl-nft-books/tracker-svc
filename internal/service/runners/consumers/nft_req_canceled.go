package consumers

import (
	"context"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

func (c *MarketPlaceConsumer) ConsumeNFTRequestCanceledEvent(ch <-chan etherdata.NFTRequestCanceledEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+nftRequestCanceledConsumerSuffix,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": c.network.FactoryAddress}

					if err = c.database.NftRequests().New().
						UpdateStatus(resources.NftRequestCanceled).
						UpdateLastUpdated(event.Timestamp).
						FilterUpdateById(event.RequestId.Int64()).Update(); err != nil {
						return errors.Wrap(err, "failed to consume nft request created transaction", logField)
					}

					// Updating contract`s last nft request created block
					if err = c.database.Blocks().UpdateNFTRequestCanceledBlockColumn(event.BlockNumber, c.network.ChainId); err != nil {
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
