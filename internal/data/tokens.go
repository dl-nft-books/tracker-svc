package data

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Token struct {
	Id        int64  `db:"id" structs:"-"`
	Contract  string `db:"contract" structs:"contract"`
	Name      string `db:"name" structs:"name"`
	Symbol    string `db:"symbol" structs:"symbol"`
	LastBlock uint64 `db:"last_block" structs:"last_block"`
}

//go:generate mockery --case=underscore --name=KeyValueQ
type TokensQ interface {
	New() TokensQ
	Page(page pgdb.OffsetPageParams) TokensQ
	Get(id int64) (*Token, error)
	Insert(token Token) (int64, error)
	UpdateLastBlock(lastBlock uint64, id int64) error
}

func (t *Token) Address() common.Address {
	return common.HexToAddress(t.Contract)
}
