package consumers

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/book-svc/connector/models"
	bookResources "gitlab.com/tokend/nft-books/book-svc/resources"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/key_value"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
)

type FactoryConsumer struct {
	logger   *logan.Entry
	cfg      config.FactoryTracker
	ctx      context.Context
	booker   *booker.Connector
	database data.DB
}

func NewFactoryConsumer(cfg config.Config, ctx context.Context) *FactoryConsumer {
	return &FactoryConsumer{
		logger:   cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.FactoryTracker(),
		booker:   cfg.BookerConnector(),
		database: postgres.NewDB(cfg.DB()),
	}
}

func (c *FactoryConsumer) ConsumeDeployedEvents(internalChannel <-chan etherdata.ContractDeployedEvent, combinersChannel chan<- common.Address) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) error {
			for {
				select {
				case event := <-internalChannel:
					if err := c.processDeployEvent(event); err != nil {
						return errors.Wrap(err, "failed to process deploy event")
					}

					entry, err := c.database.Contracts().GetByContract(event.Address.String())
					if err != nil {
						return errors.Wrap(err, "failed to validate whether the contract already exists")
					}
					if entry != nil {
						c.logger.Warnf("Contract with address %s already exists. Omitting\n", event.Address.String())
						continue
					}

					_, err = c.database.Contracts().Insert(data.Contract{
						Contract:  event.Address.String(),
						Name:      event.Name,
						Symbol:    event.Symbol,
						LastBlock: event.BlockNumber,
					})
					if err != nil {
						return errors.Wrap(err, "failed to insert contract into the database")
					}

					combinersChannel <- event.Address
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *FactoryConsumer) processDeployEvent(event etherdata.ContractDeployedEvent) error {
	eventTokenId := int64(event.TokenId)
	bookResponse, err := c.booker.ListBooks(models.ListBooksParams{
		TokenId: []int64{eventTokenId},
	})
	if err != nil {
		return errors.Wrap(err, "failed to call a list books function from booker", logan.F{
			"token_id": eventTokenId,
		})
	}
	if len(bookResponse.Data) == 0 {
		c.logger.WithFields(logan.F{"token_id": eventTokenId}).Warnf("book with specified token id was not found")
		return nil
	}
	var (
		book   = bookResponse.Data[0]
		bookId = cast.ToInt64(book.Key.ID)
	)

	switch event.Status {
	case types.ReceiptStatusSuccessful:
		var (
			contractAddress     = event.Address.String()
			successDeployStatus = bookResources.DeploySuccessful
		)

		if err = c.booker.UpdateBook(models.UpdateBookParams{
			Id:              bookId,
			ContractAddress: &contractAddress,
			DeployStatus:    &successDeployStatus,
		}); err != nil {
			return errors.Wrap(err, "failed to update book via connector")
		}

		return c.database.KeyValue().Upsert(data.KeyValue{
			Key:   key_value.FactoryTrackerCursor,
			Value: strconv.FormatUint(event.BlockNumber, 10),
		})
	case types.ReceiptStatusFailed:
		var deployFailedStatus = bookResources.DeployFailed

		return c.booker.UpdateBook(models.UpdateBookParams{
			Id:           bookId,
			DeployStatus: &deployFailedStatus,
		})
	}

	return errors.From(errors.New(""), logan.F{
		"block_number": event.BlockNumber,
	})
}
