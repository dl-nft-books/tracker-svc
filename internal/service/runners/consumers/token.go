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
	"gitlab.com/tokend/nft-books/network-svc/connector/models"
)

const (
	transferConsumerSuffix      = "-token-transfer"
	mintConsumerSuffix          = "-token-mint"
	mintByNftConsumerSuffix     = "-token-mint-by-nft"
	updateConsumerSuffix        = "-token-update"
	updateVoucherConsumerSuffix = "-voucher-update"
)

type TokenConsumer struct {
	logger   *logan.Entry
	cfg      config.Runner
	ctx      context.Context
	database data.DB

	ipfsLoader  *helpers.IpfsLoader
	network     models.NetworkDetailedResponse
	booker      *booker.Connector
	generatorer *generatorer.Connector
	documenter  *documenter.Connector
}

func NewTokenConsumer(cfg config.Config, ctx context.Context, network models.NetworkDetailedResponse) *TokenConsumer {
	return &TokenConsumer{
		logger:     cfg.Log(),
		ctx:        ctx,
		cfg:        cfg.Consumers(),
		database:   postgres.NewDB(cfg.DB()),
		network:    network,
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

					// Getting book info by task id
					book, err := c.booker.GetBookById(task.Attributes.BookId)
					if err != nil {
						return errors.Wrap(err, "failed to get book of specified task", logField.Merge(
							logan.F{"task_id": task.ID, "book_id": task.Attributes.BookId}))
					}
					if book == nil {
						c.logger.
							WithFields(logField.Merge(logan.F{"book_id": task.Attributes.BookId})).
							Warn("could not find book")
						continue
					}
					if err = c.database.Transaction(func() error {
						if err = c.UpdatingTransaction(address, book, task); err != nil {
							return err
						}
						if err = c.MintUpdating(address, book.Data.Attributes.ChainId, task, event); err != nil {
							return err
						}
						return nil
					}); err != nil {
						return errors.Wrap(err, "failed to create new token or token is already exists", logField)
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

func (c *TokenConsumer) ConsumeMintByNftEvents(address common.Address, ch <-chan etherdata.SuccessfullyMintedByNftEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+mintByNftConsumerSuffix,
		func(ctx context.Context) (err error) {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": address.String()}

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

					if err = c.database.Transaction(func() error {
						if err = c.UpdatingTransaction(address, book, task); err != nil {
							return err
						}
						if err = c.MintNftUpdating(address, book.Data.Attributes.ChainId, task, event); err != nil {
							return err
						}
						return nil
					}); err != nil {
						return errors.Wrap(err, "failed to create new token or token is already exists", logField)
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

func (c *TokenConsumer) CheckPayment(bookUrl string, f logan.F) error {
	check, err := c.database.Payments().New().FilterByBookUrl(bookUrl).Get()

	if err != nil {
		return errors.Wrap(err, "failed to check is payment exist", f)
	}
	if check != nil {
		c.logger.WithFields(logan.F{"book_url": bookUrl}).Warn("payment with such book_url is already exist")
		return errors.New("payment with such book_url is already exist")
	}
	return nil
}

func (c *TokenConsumer) UpdatingTransaction(
	address common.Address,
	book *bookerModels.GetBookResponse,
	task generatorerResources.Task) error {
	logField := logan.F{"contract_address": address.String()}
	// Getting nft banner img link
	bannerLink, err := c.documenter.GetDocumentLink(book.Data.Attributes.Banner.Attributes.Key)
	if err != nil {
		return errors.Wrap(err, "failed to get banner image link", logField)
	}
	// Uploading metadata
	if err = c.ipfsLoader.UploadMetadata(opensea.Metadata{
		Name:        fmt.Sprintf("%s #%s", book.Data.Attributes.Title, book.Data.Attributes.TokenId),
		Description: book.Data.Attributes.Description,
		Image:       bannerLink.Data.Attributes.Url,
		FileURL:     c.ipfsLoader.BaseUri + task.Attributes.FileIpfsHash,
	}); err != nil {
		return errors.Wrap(err, "failed to load metadata to the ipfs")
	}

	// Uploading file
	if err = c.ipfsLoader.UploadFile(task.Attributes.FileIpfsHash); err != nil {
		return errors.Wrap(err, "failed to load file to the ipfs", logField)
	}

	// Updating task info
	taskStatus := generatorerResources.TaskFinishedUploading
	if err = c.generatorer.UpdateTask(generatorerModels.UpdateTaskParams{
		Id:      cast.ToInt64(task.ID),
		Status:  &taskStatus,
		TokenId: &book.Data.Attributes.TokenId,
	}); err != nil {
		return errors.Wrap(err, "failed to update task`s token id and status", logField)
	}

	return nil
}

func (c *TokenConsumer) MintUpdating(
	address common.Address,
	chainId int64,
	task generatorerResources.Task,
	event etherdata.SuccessfulMintEvent) error {
	logField := logan.F{"contract_address": address.String()}
	// Getting contract by address
	contract, err := c.database.Contracts().GetByAddress(address.String())
	if err != nil {
		return errors.Wrap(err, "failed to update status", logField)
	}
	//Check if Payment with such book_url is already exists
	if err := c.CheckPayment(c.ipfsLoader.BaseUri+task.Attributes.FileIpfsHash, logField); err != nil {
		return err
	}
	// Inserting information about payment
	paymentId, err := c.database.Payments().New().Insert(data.Payment{
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
		BookUrl:           c.ipfsLoader.BaseUri + task.Attributes.FileIpfsHash,
	})
	if err != nil {
		return errors.Wrap(err, "failed to add payment to the table", logField)
	}
	// Inserting information about token
	if _, err = c.generatorer.CreateToken(generatorerModels.CreateTokenParams{
		Account:        event.Recipient.String(),
		MetadataHash:   task.Attributes.MetadataIpfsHash,
		Status:         generatorerResources.TokenFinishedUploading,
		TokenId:        event.TokenId,
		Signature:      task.Attributes.Signature,
		BookId:         task.Attributes.BookId,
		PaymentId:      paymentId,
		ChainId:        chainId,
		IsTokenPayment: true,
	}); err != nil {
		return errors.Wrap(err, "failed to create new token or token is already exists", logField)
	}

	// Updating contract`s last mint block
	if err = c.database.Contracts().UpdatePreviousMintBlock(event.BlockNumber).Update(contract.Id); err != nil {
		return errors.Wrap(err, "failed to update contract`s last mint block", logField.Merge(logan.F{
			"contract_id": contract.Id,
		}))
	}
	return nil
}

func (c *TokenConsumer) MintNftUpdating(
	address common.Address,
	chainId int64,
	task generatorerResources.Task,
	event etherdata.SuccessfullyMintedByNftEvent) error {
	logField := logan.F{"contract_address": address.String()}
	// Getting contract by address
	contract, err := c.database.Contracts().GetByAddress(address.String())
	if err != nil {
		return errors.Wrap(err, "failed to update status", logField)
	}

	//Check if Payment with such book_url is already exists
	check, err := c.database.NftPayments().New().FilterByBookUrl(c.ipfsLoader.BaseUri + task.Attributes.FileIpfsHash).Get()

	if err != nil {
		return errors.Wrap(err, "failed to check is payment exist", logField)
	}
	if check != nil {
		c.logger.WithFields(logan.F{"book_url": c.ipfsLoader.BaseUri + task.Attributes.FileIpfsHash}).Warn("payment with such book_url is already exist")
		return errors.New("payment with such book_url is already exist")
	}

	// Inserting information about payment
	paymentId, err := c.database.NftPayments().New().Insert(data.NftPayment{
		ContractId:        contract.Id,
		ContractAddress:   contract.Addr,
		PayerAddress:      event.Recipient.String(),
		NftAddress:        event.NftAddress.String(),
		NftId:             event.NftId,
		FloorPrice:        event.NftFloorPrice.String(),
		PriceMinted:       event.MintedTokenPrice.String(),
		PurchaseTimestamp: event.Timestamp,
		BookUrl:           c.ipfsLoader.BaseUri + task.Attributes.FileIpfsHash,
	})
	if err != nil {
		return errors.Wrap(err, "failed to add payment to the table", logField)
	}
	// Inserting information about token
	if _, err = c.generatorer.CreateToken(generatorerModels.CreateTokenParams{
		Account:        event.Recipient.String(),
		MetadataHash:   task.Attributes.MetadataIpfsHash,
		Status:         generatorerResources.TokenFinishedUploading,
		TokenId:        event.TokenId,
		Signature:      task.Attributes.Signature,
		BookId:         task.Attributes.BookId,
		PaymentId:      paymentId,
		ChainId:        chainId,
		IsTokenPayment: false,
	}); err != nil {
		return errors.Wrap(err, "failed to create new token or token is already exists", logField)
	}

	// Updating contract`s last mint block
	if err = c.database.Contracts().UpdatePreviousMintByNftBlock(event.BlockNumber).Update(contract.Id); err != nil {
		return errors.Wrap(err, "failed to update contract`s last mint by nft block", logField.Merge(logan.F{
			"contract_id": contract.Id,
		}))
	}
	return nil
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

					contract, err := c.database.Contracts().GetByAddress(address.String())
					lastUpdateBlock, err := c.database.Blocks().FilterByContractId(contract.Id).Get()
					if err != nil {
						return errors.Wrap(err, "failed to get last block to the given address", logField)
					}
					if lastUpdateBlock.TransferBlock > event.BlockNumber {
						c.logger.WithFields(logField).Warnf("Event has been already consumed with a block number of %d", event.BlockNumber)
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

						if err = c.database.Blocks().UpdateTransferBlock(event.BlockNumber, contract.Id); err != nil {
							return errors.Wrap(err, "failed to update block")
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
					contract, err := c.database.Contracts().GetByAddress(address.String())
					lastUpdateBlock, err := c.database.Blocks().FilterByContractId(contract.Id).Get()
					if err != nil {
						return errors.Wrap(err, "failed to get last block to the given address", logField)
					}
					if lastUpdateBlock.UpdateBlock > event.BlockNumber {
						c.logger.WithFields(logField).Warnf("Event has been already consumed with a block number of %d", event.BlockNumber)
						continue
					}
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
						Id:           bookId,
						Title:        &event.Name,
						Symbol:       &event.Symbol,
						Price:        &event.Price,
						ContractName: &event.Name,
						FloorPrice:   &event.FloorPrice,
					}); err != nil {
						return errors.Wrap(err, "failed to update book parameters")
					}

					if err != nil {
						return errors.Wrap(err, "failed to get contract")
					}
					if contract.Name != event.Name || contract.Symbol != event.Symbol {
						contractsQ := c.database.Contracts()
						if contract.Name != event.Name {
							contractsQ = contractsQ.UpdateName(event.Name)
						}
						if contract.Symbol != event.Symbol {
							contractsQ = contractsQ.UpdateSymbol(event.Symbol)
						}
						if err = contractsQ.Update(contract.Id); err != nil {
							return errors.Wrap(err, "failed to update contract parameters")
						}

					}
					if err = c.database.Blocks().UpdateParamsUpdateBlock(event.BlockNumber, contract.Id); err != nil {
						return errors.Wrap(err, "failed to update block")
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

func (c *TokenConsumer) ConsumeVoucherUpdateEvents(address common.Address, ch <-chan etherdata.VoucherUpdateEvent) {
	running.WithBackOff(
		c.ctx,
		c.logger,
		c.cfg.Prefix+updateVoucherConsumerSuffix,
		func(ctx context.Context) error {
			for {
				select {
				case event := <-ch:
					logField := logan.F{"contract_address": address.String()}

					contract, err := c.database.Contracts().GetByAddress(address.String())
					lastUpdateBlock, err := c.database.Blocks().FilterByContractId(contract.Id).Get()
					if err != nil {
						return errors.Wrap(err, "failed to get last block to the given address", logField)
					}
					if lastUpdateBlock.VoucherUpdateBlock > event.BlockNumber {
						c.logger.WithFields(logField).Warnf("Event has been already consumed with a block number of %d", event.BlockNumber)
						continue
					}
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

					var (
						bookId             = cast.ToInt64(bookResponse.Data[0].Key.ID)
						voucherToken       = event.VoucherTokenAddress.String()
						voucherTokenAmount = event.VoucherTokenAmount.String()
					)

					if err = c.booker.UpdateBook(bookerModels.UpdateBookParams{
						Id:                 bookId,
						VoucherToken:       &voucherToken,
						VoucherTokenAmount: &voucherTokenAmount,
					}); err != nil {
						return errors.Wrap(err, "failed to update book parameters")
					}

					if err != nil {
						return errors.Wrap(err, "failed to get contract")
					}
					if err = c.database.Blocks().UpdateParamsVoucherUpdateBlock(event.BlockNumber, contract.Id); err != nil {
						return errors.Wrap(err, "failed to update block")
					}

					c.logger.WithFields(logField).Infof("Successfully processed voucher update event with a block number of %d", event.BlockNumber)
				}
			}
		},
		c.cfg.Backoff.NormalPeriod,
		c.cfg.Backoff.MinAbnormalPeriod,
		c.cfg.Backoff.MaxAbnormalPeriod,
	)
}
