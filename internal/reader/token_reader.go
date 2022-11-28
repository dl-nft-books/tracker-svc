package reader

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

type TokenReader interface {
	From(from uint64) TokenReader
	To(to uint64) TokenReader
	WithAddress(address common.Address) TokenReader
	WithCtx(ctx context.Context) TokenReader
	WithRPC(rpc *ethclient.Client) TokenReader

	GetRPCInstance(chainID int64) (*ethclient.Client, error)
	GetSuccessfulMintEvents(chainID int64) ([]ethereum.SuccessfulMintEvent, error)
	GetTransferEvents() ([]ethereum.TransferEvent, error)
}
