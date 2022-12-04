package processors

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	generatorer "gitlab.com/tokend/nft-books/generator-svc/connector/api"
)

type tokenProcessor struct {
	logger      *logan.Entry
	booker      *booker.Connector
	generatorer *generatorer.Connector
	database    data.DB
}

func NewTokenConsumer(cfg config.Config, ctx context.Context) ethereum.TokenProcessor {
	return &tokenProcessor{
		logger:      cfg.Log(),
		booker:      cfg.BookerConnector(),
		generatorer: cfg.GeneratorConnector(),
		database:    postgres.NewDB(cfg.DB()),
	}
}

func (p *tokenProcessor) ProcessSuccessfulMint(event etherdata.SuccessfulMintEvent) error {
	//TODO implement me
	panic("implement me")
}

func (p *tokenProcessor) ProcessTransfer(event etherdata.TransferEvent) error {
	//TODO implement me
	panic("implement me")
}

func (p *tokenProcessor) ProcessUpdate(event etherdata.UpdateEvent) error {
	//TODO implement me
	panic("implement me")
}
