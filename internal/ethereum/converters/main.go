package converters

import (
	"context"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/tracker-svc/solidity/generated/erc20"
	"github.com/dl-nft-books/tracker-svc/solidity/generated/marketplace"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"time"
)

type EventConverter struct {
	client      *ethclient.Client
	ctx         context.Context
	nativeToken etherdata.Erc20Info
}

func NewEventConverter(client *ethclient.Client, ctx context.Context, nativeToken etherdata.Erc20Info) EventConverter {
	return EventConverter{
		client:      client,
		ctx:         ctx,
		nativeToken: nativeToken,
	}
}

func (c *EventConverter) TokenSuccessfullyPurchased(raw marketplace.MarketplaceTokenSuccessfullyPurchased) (*etherdata.TokenSuccessfullyPurchasedEvent, error) {
	receipt, err := c.client.TransactionReceipt(c.ctx, raw.Raw.TxHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
			"tx_hash": raw.Raw.TxHash.String(),
		})
	}

	purchaseTimestamp, err := c.getBlockTimestamp(raw.Raw.BlockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block timestamp")
	}

	erc20Data, err := c.GetErc20Data(raw.BuyParams.TokenContract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get erc20 data from the contract")
	}

	return &etherdata.TokenSuccessfullyPurchasedEvent{
		ContractAddress:   raw.Raw.Address,
		Recipient:         raw.Recipient,
		TokenId:           raw.BuyParams.FutureTokenId.Int64(),
		Uri:               raw.BuyParams.TokenURI,
		MintedTokenPrice:  raw.MintedTokenPrice,
		Erc20Info:         *erc20Data,
		Amount:            raw.PaidTokensAmount,
		PaymentTokenPrice: raw.BuyParams.PaymentDetails.PaymentTokenPrice,
		Status:            receipt.Status,
		BlockNumber:       raw.Raw.BlockNumber,
		Timestamp:         *purchaseTimestamp,
		Type:              raw.PaymentType,
	}, nil
}

//
//func (c *EventConverter) SuccessfulMintByNft(raw marketplace.MarketplaceSuccessfullyMintedByNFT) (*etherdata.SuccessfullyMintedByNftEvent, error) {
//	receipt, err := c.client.TransactionReceipt(c.ctx, raw.Raw.TxHash)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
//			"tx_hash": raw.Raw.TxHash.String(),
//		})
//	}
//
//	purchaseTimestamp, err := c.getBlockTimestamp(raw.Raw.BlockNumber)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to get block timestamp")
//	}
//
//	return &etherdata.SuccessfullyMintedByNftEvent{
//		Recipient:        raw.Recipient,
//		TokenId:          raw.MintedTokenInfo.TokenId.Int64(),
//		Uri:              raw.MintedTokenInfo.TokenURI,
//		MintedTokenPrice: raw.MintedTokenInfo.MintedTokenPrice,
//		NftFloorPrice:    raw.NftFloorPrice,
//		NftAddress:       raw.NftAddress,
//		NftId:            raw.TokenId.Int64(),
//		Status:           receipt.Status,
//		BlockNumber:      raw.Raw.BlockNumber,
//		Timestamp:        *purchaseTimestamp,
//	}, nil
//}
//
////func (c *EventConverter) UpdateVoucher(raw marketplace.MarketplaceTokenContractParamsUpdated) etherdata.VoucherUpdateEvent {
////	return etherdata.VoucherUpdateEvent{
////		VoucherTokenAddress: raw.NewVoucherTokenContract,
////		VoucherTokenAmount:  raw.NewVoucherTokensAmount,
////		BlockNumber:         raw.Raw.BlockNumber,
////	}
////}

// getBlockTimestamp is a function that returns a timestamp
// when a block with specified blockNumber was initialized
func (c *EventConverter) getBlockTimestamp(blockNumber uint64) (*time.Time, error) {
	// Get header by a block number
	header, err := c.client.BlockByNumber(c.ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block header by its number")
	}

	// Get timestamp from a block as an uint64 and convert it to time.Time
	blockTime := time.Unix(int64(header.Time()), 0)
	return &blockTime, nil
}

// GetErc20Data is a function that retrieves the etherdata.Erc20Info from the
// address specified in parameters. If address is etherdata.NullAddress, function
// returns the native erc20 data. Otherwise, contract on the specified
// address must implement IERC20 interface.
func (c *EventConverter) GetErc20Data(address common.Address) (*etherdata.Erc20Info, error) {
	if address == etherdata.NullAddress {
		return &c.nativeToken, nil
	}

	erc20Instance, err := erc20.NewErc20(address, c.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize erc20 instance", logan.F{
			"erc20_address": address,
		})
	}

	tokenName, err := erc20Instance.Name(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to read marketplace's name from the contract instance")
	}

	tokenSymbol, err := erc20Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to read marketplace's symbol from the contract instance")
	}

	tokenDecimals, err := erc20Instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get marketplace's decimals")
	}

	return &etherdata.Erc20Info{
		TokenAddress: address,
		Name:         tokenName,
		Symbol:       tokenSymbol,
		Decimals:     tokenDecimals,
	}, nil
}
