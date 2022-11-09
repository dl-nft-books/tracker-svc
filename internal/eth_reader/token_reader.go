package eth_reader

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/erc20"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/itokencontract"
	"math/big"
	"time"
)

var (
	NullIteratorErr = errors.New("iterator has a nil value")

	NullAddress      = common.Address{}
	DefaultErc20Info = Erc20Info{
		TokenAddress: common.Address{},
		Name:         "Ethereum",
		Symbol:       "ETH",
		Decimals:     18,
	}
)

type TokenContractReader struct {
	rpc *ethclient.Client
}

func NewTokenContractReader(rpc *ethclient.Client) TokenContractReader {
	return TokenContractReader{
		rpc: rpc,
	}
}

type Erc20Info struct {
	TokenAddress common.Address
	Name         string
	Symbol       string
	Decimals     uint8
}

type SuccessfulMintEvent struct {
	Recipient         common.Address
	TokenId           int64
	Uri               string
	Erc20Info         Erc20Info
	Amount            *big.Int
	PaymentTokenPrice *big.Int
	MintedTokenPrice  *big.Int
	Status            uint64
	BlockNumber       uint64
	Timestamp         time.Time
}

func (r *TokenContractReader) GetSuccessfulMintEvents(
	contract common.Address,
	startBlock,
	endBlock uint64,
) (
	events []SuccessfulMintEvent,
	lastBlock uint64,
	err error,
) {
	instance, err := itokencontract.NewItokencontract(contract, r.rpc)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterSuccessfullyMinted(
		&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil, nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, 0, errors.From(NullIteratorErr, logan.F{
			"contract": contract.String(),
		})
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event

		if event != nil {
			receipt, err := r.rpc.TransactionReceipt(context.Background(), event.Raw.TxHash)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			purchaseTimestamp, err := r.getBlockTimestamp(event.Raw.BlockNumber)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to get block timestamp")
			}

			if event.PaymentTokenAddress == NullAddress {
				events = append(events, SuccessfulMintEvent{
					Recipient:         event.Recipient,
					TokenId:           event.MintedTokenInfo.TokenId.Int64(),
					Uri:               event.MintedTokenInfo.TokenURI,
					MintedTokenPrice:  event.MintedTokenInfo.PricePerOneToken,
					PaymentTokenPrice: event.PaymentTokenPrice,
					Erc20Info:         DefaultErc20Info,
					Amount:            event.PaidTokensAmount,
					Status:            receipt.Status,
					BlockNumber:       event.Raw.BlockNumber,
					Timestamp:         *purchaseTimestamp,
				})

				continue
			}

			erc20Data, err := r.GetErc20Data(event.PaymentTokenAddress)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to get erc20 data from the contract")
			}

			events = append(events, SuccessfulMintEvent{
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

			lastBlock = event.Raw.BlockNumber
		}
	}

	return
}

func (r *TokenContractReader) getBlockTimestamp(blockNumber uint64) (*time.Time, error) {
	// Get header by a block number
	header, err := r.rpc.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block header by its number")
	}

	// Get timestamp from a block as an uint64 and convert it to time.Time
	blockTime := time.Unix(int64(header.Time()), 0)
	return &blockTime, nil
}

func (r *TokenContractReader) GetErc20Data(address common.Address) (*Erc20Info, error) {
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

	return &Erc20Info{
		TokenAddress: address,
		Name:         tokenName,
		Symbol:       tokenSymbol,
		Decimals:     tokenDecimals,
	}, nil
}
