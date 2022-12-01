package contract_reader

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

type TokenFactoryReader interface {
	From(from uint64) TokenFactoryReader
	To(to uint64) TokenFactoryReader
	WithAddress(address common.Address) TokenFactoryReader
	WithCtx(ctx context.Context) TokenFactoryReader

	GetContractCreatedEvents() ([]ethereum.ContractCreatedEvent, error)
}
