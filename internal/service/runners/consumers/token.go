package consumers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	bookerModels "gitlab.com/tokend/nft-books/book-svc/connector/models"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	generatorer "gitlab.com/tokend/nft-books/generator-svc/connector"
	generatorerModels "gitlab.com/tokend/nft-books/generator-svc/connector/models"
)

type TokenConsumer struct {
	logger   *logan.Entry
	cfg      config.FactoryTracker
	ctx      context.Context
	database data.DB

	booker      *booker.Connector
	generatorer *generatorer.Connector
}

func NewTokenConsumer(cfg config.Config, ctx context.Context) *TokenConsumer {
	return &TokenConsumer{
		logger:   cfg.Log(),
		ctx:      ctx,
		cfg:      cfg.FactoryTracker(),
		database: postgres.NewDB(cfg.DB()),

		booker:      cfg.BookerConnector(),
		generatorer: cfg.GeneratorConnector(),
	}
}

func (c *TokenConsumer) ConsumeMintEvents(address common.Address, ch <-chan etherdata.SuccessfulMintEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					/*
						return t.trackerDB.Transaction(func() error {
								// FIXME: Make the following actions via connectors:
								// 1. Get task info using event uri
								// 2. Get book info using retrieved from step 1 (e.g., via included)
								// 3. Insert new token into generator table
								// 4. Update status of a task

								// Updating book and tasks db
								t.booksQ = t.booksQ.New()

								// Getting task by hash (uri)
								task, err := t.generatorDB.Tasks().GetByHash(event.Uri)
								if err != nil {
									return errors.Wrap(err, "failed to get task by ipfs hash")
								}
								if task == nil {
									t.log.Warnf("Could not find task with a hash of %s", event.Uri)
									return nil
								}

								// Getting book info
								book, err := t.booksQ.FilterActual().FilterByID(task.BookId).Get()
								if err != nil {
									return errors.Wrap(err, "failed to get book by task id")
								}
								if book == nil {
									t.log.WithFields(logan.F{"book_id": task.BookId}).Warn("Could not find book")
									return nil
								}

								// Parsing banner key
								bannerKey, err := helpers.GetDocumentKey(book.Banner)
								if err != nil {
									return errors.Wrap(err, "failed to get document key")
								}
								if bannerKey == nil {
									return bannerNotFoundErr
								}

								// Getting nft banner img link
								bannerLink, err := t.documenter.GetDocumentLink(*bannerKey)
								if err != nil {
									return errors.Wrap(err, "failed to get banner image link")
								}

								// updating status to loading on IPFS
								if err = t.generatorDB.Tasks().UpdateStatus(resources.TaskUploading, task.Id); err != nil {
									return errors.Wrap(err, "failed to update status", logan.F{
										"task_id": task.Id,
									})
								}

								// inserting information about payment
								paymentId, err := t.trackerDB.Payments().Insert(data.Payment{
									ContractId:        contract.Id,
									ContractAddress:   contract.Contract,
									PayerAddress:      event.Recipient.String(),
									TokenAddress:      event.Erc20Info.TokenAddress.String(),
									TokenSymbol:       event.Erc20Info.Symbol,
									TokenName:         event.Erc20Info.Name,
									TokenDecimals:     event.Erc20Info.Decimals,
									Amount:            event.Amount.String(),
									PriceMinted:       event.MintedTokenPrice.String(),
									PriceToken:        event.PaymentTokenPrice.String(),
									PurchaseTimestamp: event.Timestamp,
									BookUrl:           baseURI + task.FileIpfsHash,
								})
								if err != nil {
									return errors.Wrap(err, "failed to add payment to the table")
								}

								// Inserting information about token
								tokenId, err := t.generatorDB.Tokens().Insert(external.Token{
									Account:      event.Recipient.String(),
									TokenId:      event.TokenId,
									BookId:       book.ID,
									PaymentId:    paymentId,
									MetadataHash: task.MetadataIpfsHash,
									Signature:    task.Signature,
									Status:       resources.TokenUploading,
								})

								// Uploading metadata
								if err = t.ipfsLoader.UploadMetadata(opensea.Metadata{
									Name:        fmt.Sprintf("%s #%v", book.Title, task.Id),
									Description: book.Description,
									Image:       bannerLink.Data.Attributes.Url,
									FileURL:     baseURI + task.FileIpfsHash,
								}); err != nil {
									return errors.Wrap(err, "failed to load metadata to the ipfs")
								}

								// Uploading file
								if err = t.ipfsLoader.UploadFile(task.FileIpfsHash); err != nil {
									return errors.Wrap(err, "failed to load file to the ipfs")
								}

								// Updating tokenID
								if err = t.generatorDB.Tasks().UpdateTokenId(event.TokenId, task.Id); err != nil {
									return errors.Wrap(err, "failed to update token id")
								}

								// Updating status to fully processed task
								if err = t.generatorDB.Tasks().UpdateStatus(resources.TaskFinishedUploading, task.Id); err != nil {
									return errors.Wrap(err, "failed to update task's status", logan.F{
										"task_id": task.Id,
									})
								}

								// Updating status to a token
								if err = t.generatorDB.Tokens().UpdateStatus(resources.TokenFinishedUploading, tokenId); err != nil {
									return errors.Wrap(err, "failed to update token's status", logan.F{
										"token_id": tokenId,
									})
								}

								return nil
							})
					*/
					c.logger.Infof("Hey, I found this nice little thing over here: %v", event)
				}
			}
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeTransferEvents(address common.Address, ch <-chan etherdata.TransferEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
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

					var (
						dbTokenId = cast.ToInt64(tokenResponse.Data[0].Key.ID)
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
		},
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}

func (c *TokenConsumer) ConsumeUpdateEvents(address common.Address, ch <-chan etherdata.UpdateEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Name,
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
		c.cfg.Runner.NormalPeriod,
		c.cfg.Runner.MinAbnormalPeriod,
		c.cfg.Runner.MaxAbnormalPeriod,
	)
}
