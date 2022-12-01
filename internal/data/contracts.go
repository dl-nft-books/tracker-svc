package data

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Contract struct {
	Id        int64  `db:"id" structs:"-"`
	Contract  string `db:"contract" structs:"contract"`
	Name      string `db:"name" structs:"name"`
	Symbol    string `db:"symbol" structs:"symbol"`
	LastBlock uint64 `db:"last_block" structs:"last_block"`
}

//go:generate mockery --case=underscore --name=KeyValueQ
type ContractsQ interface {
	New() ContractsQ
	Select() ([]Contract, error)
	Page(page pgdb.OffsetPageParams) ContractsQ
	Get(id int64) (*Contract, error)
	Insert(contract ...Contract) ([]int64, error)
	UpdateLastBlock(lastBlock uint64, id int64) error
}

func (c *Contract) Address() common.Address {
	return common.HexToAddress(c.Contract)
}
