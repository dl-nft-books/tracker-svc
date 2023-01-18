package data

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Contract struct {
	Id                int64  `db:"id" structs:"-"`
	Addr              string `db:"address" structs:"address"`
	Name              string `db:"name" structs:"name"`
	Symbol            string `db:"symbol" structs:"symbol"`
	PreviousMintBLock uint64 `db:"previous_mint_block" structs:"previous_mint_block"`
	ChainId           int64  `db:"chain_id" structs:"chain_id"`
}

//go:generate mockery --case=underscore --name=KeyValueQ
type ContractsQ interface {
	New() ContractsQ
	Select() ([]Contract, error)
	Page(page pgdb.OffsetPageParams) ContractsQ
	Get(id int64) (*Contract, error)
	GetByAddress(address string) (*Contract, error)
	FilterByChainId(chainId int64) ContractsQ
	Insert(contract Contract) (int64, error)
	UpdatePreviousMintBlock(lastBlock uint64, id int64) error
}

func (c *Contract) Address() common.Address {
	return common.HexToAddress(c.Addr)
}
