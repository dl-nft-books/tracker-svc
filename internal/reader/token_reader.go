package reader

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

type TokenReader interface {
	From(from uint64) TokenReader
	To(to uint64) TokenReader
	WithAddress(address common.Address) TokenReader

	GetSuccessfulMintEvents() ([]ethereum.SuccessfulMintEvent, error)
	GetTransferEvents() ([]ethereum.TransferEvent, error)
}
