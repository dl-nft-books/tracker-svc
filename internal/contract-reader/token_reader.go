package contract_reader

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum/token"
)

type TokenReader interface {
	From(from uint64) TokenReader
	To(to uint64) TokenReader
	WithAddress(address common.Address) TokenReader
	WithCtx(ctx context.Context) TokenReader

	GetSuccessfulMintEvents() ([]token.SuccessfulMintEvent, error)
	GetTransferEvents() ([]token.TransferEvent, error)
	GetUpdateEvents() ([]token.UpdateEvent, error)
}
