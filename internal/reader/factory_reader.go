package reader

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

type FactoryReader interface {
	From(from uint64) FactoryReader
	To(to uint64) FactoryReader
	WithAddress(address common.Address) FactoryReader
	WithCtx(ctx context.Context) FactoryReader

	GetContractCreatedEvents() ([]ethereum.ContractCreatedEvent, error)
}
