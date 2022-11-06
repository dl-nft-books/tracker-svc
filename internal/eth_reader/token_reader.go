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
)

var NullIteratorErr = errors.New("iterator has a nil value")

type TokenContractReader struct {
	rpc *ethclient.Client
}

func NewTokenContractReader(rpc *ethclient.Client) TokenContractReader {
	return TokenContractReader{
		rpc: rpc,
	}
}

type TokenMintedEvent struct {
	Recipient           common.Address
	BlockNumber, Status uint64
	TokenId             int64
	Uri                 string
}

func (r *TokenContractReader) GetMintEvents(
	contract common.Address,
	startBlock,
	endBlock uint64,
) (
	events []TokenMintedEvent,
	lastBlock uint64,
	err error,
) {
	instance, err := itokencontract.NewItokencontract(contract, r.rpc)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterTokenMinted(
		&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil,
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

			events = append(events, TokenMintedEvent{
				Recipient:   event.Recipient,
				BlockNumber: event.Raw.BlockNumber,
				TokenId:     event.TokenId.Int64(),
				Uri:         event.TokenURI,
				Status:      receipt.Status,
			})

			lastBlock = event.Raw.BlockNumber
		}
	}

	return
}

type TokenPaymentEvent struct {
	PayerAddress, TokenAddress  common.Address
	Amount, Price, Name, Symbol string
	Status, BlockNumber         uint64
}

func (r *TokenContractReader) GetPaymentEvents(
	contract common.Address,
	startBlock,
	endBlock uint64,
) (
	events []TokenPaymentEvent,
	lastBlock uint64,
	err error,
) {
	instance, err := itokencontract.NewItokencontract(contract, r.rpc)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterPaymentSuccessful(
		&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil,
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

			erc20Instance, err := erc20.NewErc20(event.TokenAddress, r.rpc)
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to initialize erc20 instance", logan.F{
					"erc20_address": event.TokenAddress,
				})
			}

			tokenName, err := erc20Instance.Name(&bind.CallOpts{})
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to read token's name from the contract instance")
			}

			tokenSymbol, err := erc20Instance.Symbol(&bind.CallOpts{})
			if err != nil {
				return nil, 0, errors.Wrap(err, "failed to read token's symbol from the contract instance")
			}

			events = append(events, TokenPaymentEvent{
				PayerAddress: common.Address{}, //TODO: When event will have recipient, add it
				TokenAddress: event.TokenAddress,
				Amount:       event.TokenAmount.String(),
				Price:        event.TokenPrice.String(),
				Status:       receipt.Status,
				Name:         tokenName,
				Symbol:       tokenSymbol,
				BlockNumber:  event.Raw.BlockNumber,
			})

			lastBlock = event.Raw.BlockNumber
		}
	}

	return
}
