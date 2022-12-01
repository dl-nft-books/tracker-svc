package evm_based_reader

import (
	"context"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/contract-reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/erc20"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokencontract"
)

var (
	NullIteratorErr = errors.New("iterator has a nil value")
)

type TokenContractReader struct {
	rpc             *ethclient.Client
	nativeTokenInfo ethereum.Erc20Info

	from    *uint64
	to      *uint64
	address *common.Address
	ctx     context.Context

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*tokencontract.Tokencontract
}

func NewTokenContractReader(cfg config.Config) contract_reader.TokenReader {
	return &TokenContractReader{
		rpc:                    cfg.EtherClient().Rpc,
		ctx:                    context.Background(),
		contractInstancesCache: make(map[common.Address]*tokencontract.Tokencontract),
	}
}

func (r *TokenContractReader) From(from uint64) contract_reader.TokenReader {
	r.from = &from
	return r
}

func (r *TokenContractReader) To(to uint64) contract_reader.TokenReader {
	r.to = &to
	return r
}

func (r *TokenContractReader) WithAddress(address common.Address) contract_reader.TokenReader {
	r.address = &address
	return r
}

func (r *TokenContractReader) WithCtx(ctx context.Context) contract_reader.TokenReader {
	r.ctx = ctx
	return r
}

func (r *TokenContractReader) validateParameters() error {
	if r.from == nil {
		return contract_reader.FromNotSpecifiedErr
	}
	if r.address == nil {
		return contract_reader.AddressNotSpecifiedErr
	}

	return nil
}

// GetSuccessfulMintEvents returns the successful mint events
// in the form of ethereum.SuccessfulMintEvent array
// based on contract, start and end blocks to search through
func (r *TokenContractReader) GetSuccessfulMintEvents() (events []ethereum.SuccessfulMintEvent, err error) {
	if err = r.validateParameters(); err != nil {
		return nil, err
	}

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterSuccessfullyMinted(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		}, nil, nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
		})
	}

	defer func(iterator *tokencontract.TokencontractSuccessfullyMintedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		event := iterator.Event

		if event != nil {
			receipt, err := r.rpc.TransactionReceipt(r.ctx, event.Raw.TxHash)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			purchaseTimestamp, err := r.getBlockTimestamp(event.Raw.BlockNumber)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get block timestamp")
			}

			erc20Data, err := r.GetErc20Data(event.PaymentTokenAddress)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get erc20 data from the contract")
			}

			events = append(events, ethereum.SuccessfulMintEvent{
				Recipient:         event.Recipient,
				TokenId:           event.MintedTokenInfo.TokenId.Int64(),
				Uri:               event.MintedTokenInfo.TokenURI,
				MintedTokenPrice:  event.MintedTokenInfo.PricePerOneToken,
				Erc20Info:         *erc20Data,
				Amount:            event.PaidTokensAmount,
				PaymentTokenPrice: event.PaymentTokenPrice,
				Status:            receipt.Status,
				BlockNumber:       event.Raw.BlockNumber,
				Timestamp:         *purchaseTimestamp,
			})
		}
	}

	return events, nil
}

// GetTransferEvents returns the transfer events
// in the form of ethereum.TransferEvent array
// based on contract, start and end blocks to search through
func (r *TokenContractReader) GetTransferEvents() (events []ethereum.TransferEvent, err error) {
	if err = r.validateParameters(); err != nil {
		return nil, err
	}

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create token contract instance")
	}

	iterator, err := instance.FilterTransfer(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		}, nil, nil, nil,
	)
	if iterator == nil {
		return nil, errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
		})
	}

	defer func(iterator *tokencontract.TokencontractTransferIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		event := iterator.Event

		if event != nil {
			events = append(events, ethereum.TransferEvent{
				From:    event.From,
				To:      event.To,
				TokenId: event.TokenId.Uint64(),
			})

		}
	}

	return events, nil
}

// getBlockTimestamp is a function that returns a timestamp
// when a block with specified blockNumber was initialized
func (r *TokenContractReader) getBlockTimestamp(blockNumber uint64) (*time.Time, error) {
	// Get header by a block number
	header, err := r.rpc.BlockByNumber(r.ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block header by its number")
	}

	// Get timestamp from a block as an uint64 and convert it to time.Time
	blockTime := time.Unix(int64(header.Time()), 0)
	return &blockTime, nil
}

// GetErc20Data is a function that retrieves the ethereum.Erc20Info from the
// address specified in parameters. If address is ethereum.NullAddress, function
// returns the native erc20 data. Otherwise, contract on the specified
// address must implement IERC20 interface.
func (r *TokenContractReader) GetErc20Data(address common.Address) (*ethereum.Erc20Info, error) {
	if address == ethereum.NullAddress {
		return &r.nativeTokenInfo, nil
	}

	erc20Instance, err := erc20.NewErc20(address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize erc20 instance", logan.F{
			"erc20_address": address,
		})
	}

	tokenName, err := erc20Instance.Name(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to read token's name from the contract instance")
	}

	tokenSymbol, err := erc20Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to read token's symbol from the contract instance")
	}

	tokenDecimals, err := erc20Instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token's decimals")
	}

	return &ethereum.Erc20Info{
		TokenAddress: address,
		Name:         tokenName,
		Symbol:       tokenSymbol,
		Decimals:     tokenDecimals,
	}, nil
}

func (r *TokenContractReader) getContractInstance(address common.Address) (*tokencontract.Tokencontract, error) {
	// Searching contract instance in cache, if not found -- create new and store
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := tokencontract.NewTokencontract(*r.address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
}
