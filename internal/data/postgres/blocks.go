package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blocksTable = "blocks"

	blocksIdColumn            = "id"
	blocksContractIdColumn    = "contract_id"
	blocksTransferBlockColumn = "transfer_block"
	blocksUpdateBlockColumn   = "update_block"
)

type blocksQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewBlocksQ(db *pgdb.DB) data.BlocksQ {
	return &blocksQ{
		db:       db,
		selector: squirrel.Select(fmt.Sprintf("%s.*", blocksTable)).From(blocksTable),
	}
}

func (q *blocksQ) New() data.BlocksQ {
	return NewBlocksQ(q.db.Clone())
}

func (q *blocksQ) FilterById(id ...int64) data.BlocksQ {
	q.selector = q.selector.Where(squirrel.Eq{blocksIdColumn: id})
	return q
}

func (q *blocksQ) FilterByContractId(contractId ...int64) data.BlocksQ {
	q.selector = q.selector.Where(squirrel.Eq{blocksContractIdColumn: contractId})
	return q
}

func (q *blocksQ) Get() (*data.Blocks, error) {
	var blocks data.Blocks

	err := q.db.Get(&blocks, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &blocks, err
}

func (q *blocksQ) Select() (blocks []data.Blocks, err error) {
	err = q.db.Select(&blocks, q.selector)
	return
}

func (q *blocksQ) Delete(id int64) error {
	statement := squirrel.Delete(blocksTable).Where(squirrel.Eq{blocksIdColumn: id})
	return q.db.Exec(statement)
}
