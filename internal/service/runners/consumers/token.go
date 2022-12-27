package consumers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	bookerModels "gitlab.com/tokend/nft-books/book-svc/connector/models"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/opensea"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	generatorer "gitlab.com/tokend/nft-books/generator-svc/connector"
	generatorerModels "gitlab.com/tokend/nft-books/generator-svc/connector/models"
	generatorerResources "gitlab.com/tokend/nft-books/generator-svc/resources"
	"log"
)

const (
	transferConsumerSuffix = "-token-transfer"
	mintConsumerSuffix     = "-token-mint"
	updateConsumerSuffix   = "-token-update"

	baseURI = "https://ipfs.io/ipfs/"
)

type TokenConsumer struct {
	logger   *logan.Entry
	cfg      config.Runner
	ctx      context.Context
	database data.DB

	ipfsLoader *helpers.IpfsLoader

	booker      *booker.Connector
	generatorer *generatorer.Connector
	documenter  *documenter.Connector
}

func NewTokenConsumer(cfg config.Config, ctx context.Context) *TokenConsumer {
	return &TokenConsumer{
		logger:   cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.Consumers(),
		database: postgres.NewDB(cfg.DB()),

		ipfsLoader: helpers.NewIpfsLoader(cfg),

		booker:      cfg.BookerConnector(),
		generatorer: cfg.GeneratorConnector(),
		documenter:  cfg.DocumenterConnector(),
	}
}

func (c *TokenConsumer) ConsumeMintEvents(address common.Address, ch <-chan etherdata.SuccessfulMintEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+mintConsumerSuffix,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": address.String()}

					// Validating that event was not previously already processed
					log.Println("event.Uri " + event.Uri)
					tokensResponse, err := c.generatorer.ListTokens(generatorerModels.ListTokensRequest{
						MetadataHash: []string{event.Uri},
					})
					if err != nil {
						return errors.Wrap(err, "failed to list tokens", logField.Merge(logan.F{
							"metadata_hash": event.Uri,
						}))
					}
					if len(tokensResponse.Data) > 0 {
						log.Println("tokensResponse " + tokensResponse.Data[0].ID)
						c.logger.
							WithFields(logField.Merge(logan.F{"metadata_hash": event.Uri})).
							Warn("token with specified metadata hash already exists")
						continue
					}

					// Getting task by hash (uri)
					tasksResponse, err := c.generatorer.ListTasks(generatorerModels.ListTasksRequest{IpfsHash: &event.Uri})
					if err != nil {
						return errors.Wrap(err, "failed to get task by ipfs hash", logField.Merge(logan.F{
							"ipfs_hash": event.Uri,
						}))
					}
					if len(tasksResponse.Data) == 0 {
						c.logger.
							WithFields(logField.Merge(logan.F{"ipfs_hash": event.Uri})).
							Warn("task with specified ipfs hash was not found")
						continue
					}

					task := tasksResponse.Data[0]

					// Getting book info by task id
					book, err := c.booker.GetBookById(task.Attributes.BookId)
					if err != nil {
						return errors.Wrap(err, "failed to get book of specified task", logField.Merge(
							logan.F{"task_id": tasksResponse.Data[0].ID, "book_id": task.Attributes.BookId}))
					}
					if book == nil {
						c.logger.
							WithFields(logField.Merge(logan.F{"book_id": task.Attributes.BookId})).
							Warn("could not find book")
						continue
					}

					// Getting nft banner img link
					bannerLink, err := c.documenter.GetDocumentLink(book.Data.Attributes.Banner.Attributes.Key)
					if err != nil {
						return errors.Wrap(err, "failed to get banner image link", logField)
					}

					// Updating status to loading on IPFS
					status := generatorerResources.TaskUploading
					if err = c.generatorer.UpdateTask(generatorerModels.UpdateTaskParams{
						Id:     cast.ToInt64(task.ID),
						Status: &status,
					}); err != nil {
						return errors.Wrap(err, "failed to update status", logField.Merge(logan.F{
							"task_id": task.ID,
						}))
					}

					// Getting contract by address
					contract, err := c.database.Contracts().GetByAddress(address.String())
					if err != nil {
						return errors.Wrap(err, "failed to update status", logField)
					}

					if err = c.database.Transaction(func() error {
						// Uploading metadata
						if err = c.ipfsLoader.UploadMetadata(opensea.Metadata{
							Name:        fmt.Sprintf("%s #%s", book.Data.Attributes.Title, task.ID),
							Description: book.Data.Attributes.Description,
							Image:       bannerLink.Data.Attributes.Url,
							FileURL:     baseURI + task.Attributes.FileIpfsHash,
						}); err != nil {
							return errors.Wrap(err, "failed to load metadata to the ipfs")
						}

						// Uploading file
						if err = c.ipfsLoader.UploadFile(task.Attributes.FileIpfsHash); err != nil {
							return errors.Wrap(err, "failed to load file to the ipfs", logField)
						}

						// Inserting information about payment
						paymentId, err := c.database.Payments().Insert(data.Payment{
							ContractId:        contract.Id,
							ContractAddress:   contract.Addr,
							PayerAddress:      event.Recipient.String(),
							TokenAddress:      event.Erc20Info.TokenAddress.String(),
							TokenSymbol:       event.Erc20Info.Symbol,
							TokenName:         event.Erc20Info.Name,
							TokenDecimals:     event.Erc20Info.Decimals,
							Amount:            event.Amount.String(),
							PriceMinted:       event.MintedTokenPrice.String(),
							PriceToken:        event.PaymentTokenPrice.String(),
							PurchaseTimestamp: event.Timestamp,
							BookUrl:           baseURI + task.Attributes.FileIpfsHash,
						})
						if err != nil {
							return errors.Wrap(err, "failed to add payment to the table", logField)
						}

						// Inserting information about token
						if _, err = c.generatorer.CreateToken(generatorerModels.CreateTokenParams{
							Account:      event.Recipient.String(),
							MetadataHash: task.Attributes.MetadataIpfsHash,
							Status:       generatorerResources.TokenFinishedUploading,
							TokenId:      event.TokenId,
							Signature:    task.Attributes.Signature,
							BookId:       task.Attributes.BookId,
							PaymentId:    paymentId,
						}); err != nil {
							return errors.Wrap(err, "failed to create new token", logField)
						}

						// Updating task info
						taskStatus := generatorerResources.TaskFinishedUploading
						if err = c.generatorer.UpdateTask(generatorerModels.UpdateTaskParams{
							Id:      cast.ToInt64(task.ID),
							Status:  &taskStatus,
							TokenId: &event.TokenId,
						}); err != nil {
							return errors.Wrap(err, "failed to update task`s token id and status", logField)
						}

						// Updating contract`s last mint block
						if err = c.database.Contracts().UpdatePreviousMintBlock(event.BlockNumber, contract.Id); err != nil {
							return errors.Wrap(err, "failed to update contract`s last mint block", logField.Merge(logan.F{
								"contract_id": contract.Id,
							}))
						}

						return nil
					}); err != nil {
						return errors.Wrap(err, "transaction failed")
					}

					c.logger.WithFields(logField).Infof("Successfully processed mint event of a token with id %d", event.TokenId)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeTransferEvents(address common.Address, ch <-chan etherdata.TransferEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+transferConsumerSuffix,
		func(ctx context.Context) error {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": address.String()}

					if event.From == etherdata.NullAddress || event.To == etherdata.NullAddress {
						c.logger.WithFields(logField).Info("Received transfer event with one address being null, omitting")
						continue
					}

					tokenId := int64(event.TokenId)
					tokenResponse, err := c.generatorer.ListTokens(generatorerModels.ListTokensRequest{
						TokenId: &tokenId,
					})
					if err != nil {
						return errors.Wrap(err, "failed to list tokens using connector")
					}
					if len(tokenResponse.Data) == 0 {
						c.logger.WithFields(logField.Merge(logan.F{"token_id": tokenId})).Warn("token with specified token id was not found")
					}

					for _, token := range tokenResponse.Data {
						paymentId := cast.ToInt64(token.Relationships.Payment.Data.ID)
						payment, err := c.database.Payments().FilterById(paymentId).Get()
						if err != nil {
							return errors.Wrap(err, "failed to get payment by token id")
						}
						if common.HexToAddress(payment.ContractAddress) != address {
							continue
						}

						var (
							dbTokenId = cast.ToInt64(token.Key.ID)
							newOwner  = event.To.String()
						)

						if err = c.generatorer.UpdateToken(generatorerModels.UpdateTokenParams{
							Id:    dbTokenId,
							Owner: &newOwner,
						}); err != nil {
							return errors.Wrap(err, "failed to update token using generatorer connector")
						}

						c.logger.WithFields(logField).Infof("Successfully processed transfer event of a token with id %d", event.TokenId)
					}
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeUpdateEvents(address common.Address, ch <-chan etherdata.UpdateEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+updateConsumerSuffix,
		func(ctx context.Context) error {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": address.String()}

					bookResponse, err := c.booker.ListBooks(bookerModels.ListBooksParams{
						Contract: []string{address.String()},
					})
					if err != nil {
						return errors.Wrap(err, "failed to get book corresponding to the given address", logField)
					}
					if len(bookResponse.Data) == 0 {
						c.logger.WithFields(logField).Warnf("Contract with specified address was not found")
						continue
					}

					bookId := cast.ToInt64(bookResponse.Data[0].Key.ID)
					if err = c.booker.UpdateBook(bookerModels.UpdateBookParams{
						Id:     bookId,
						Title:  &event.Name,
						Symbol: &event.Symbol,
						Price:  &event.Price,
					}); err != nil {
						return errors.Wrap(err, "failed to update book parameters")
					}

					c.logger.WithFields(logField).Infof("Successfully processed update event with a block number of %d", event.BlockNumber)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}
