package processors

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/book-svc/connector/models"
	bookResources "gitlab.com/tokend/nft-books/book-svc/resources"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/key_value"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"strconv"
)

type factoryProcessor struct {
	logger   *logan.Entry
	booker   *booker.Connector
	database data.DB
}

func NewFactoryProcessor(cfg config.Config, ctx context.Context) ethereum.FactoryProcessor {
	return &factoryProcessor{
		logger:   cfg.Log(),
		booker:   cfg.BookerConnector(),
		database: postgres.NewDB(cfg.DB()),
	}
}

func (p *factoryProcessor) ProcessDeploy(event etherdata.ContractDeployedEvent) error {
	eventTokenId := int64(event.TokenId)
	bookResponse, err := p.booker.ListBooks(models.ListBooksParams{
		TokenId: []int64{eventTokenId},
	})
	if err != nil {
		return errors.Wrap(err, "failed to call a list books function from booker", logan.F{
			"token_id": eventTokenId,
		})
	}
	if len(bookResponse.Data) == 0 {
		p.logger.WithFields(logan.F{"token_id": eventTokenId}).Warnf("book with specified token id was not found")
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

		if err = p.booker.UpdateBook(models.UpdateBookParams{
			Id:              bookId,
			ContractAddress: &contractAddress,
			DeployStatus:    &successDeployStatus,
		}); err != nil {
			return errors.Wrap(err, "failed to update book via connector")
		}

		return p.database.KeyValue().Upsert(data.KeyValue{
			Key:   key_value.FactoryTrackerCursor,
			Value: strconv.FormatUint(event.BlockNumber, 10),
		})
	case types.ReceiptStatusFailed:
		var deployFailedStatus = bookResources.DeployFailed

		return p.booker.UpdateBook(models.UpdateBookParams{
			Id:           bookId,
			DeployStatus: &deployFailedStatus,
		})
	}

	return errors.From(errors.New(""), logan.F{
		"block_number": event.BlockNumber,
	})
}
