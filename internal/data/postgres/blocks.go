package postgres

import (
	"database/sql"
	"fmt"
	"github.com/fatih/structs"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blocksTable = "blocks"

	blocksIdColumn                 = "id"
	blocksContractIdColumn         = "contract_id"
	blocksTransferBlockColumn      = "transfer_block"
	blocksUpdateBlockColumn        = "update_block"
	blocksVoucherUpdateBlockColumn = "voucher_update_block"
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

func (q *blocksQ) Insert(blocks data.Blocks) (id int64, err error) {
	statement := squirrel.Insert(blocksTable).
		Suffix("returning id").
		SetMap(structs.Map(&blocks))

	err = q.db.Get(&id, statement)
	return id, err
}

func (q *blocksQ) UpdateTransferBlock(newTransferBlock uint64, id int64) error {
	statement := squirrel.Update(blocksTable).
		Set(blocksTransferBlockColumn, newTransferBlock).
		Where(squirrel.Eq{blocksIdColumn: id})
	return q.db.Exec(statement)
}

func (q *blocksQ) UpdateParamsUpdateBlock(newUpdateBlock uint64, id int64) error {
	statement := squirrel.Update(blocksTable).
		Set(blocksUpdateBlockColumn, newUpdateBlock).
		Where(squirrel.Eq{blocksIdColumn: id})
	return q.db.Exec(statement)
}

func (q *blocksQ) UpdateParamsVoucherUpdateBlock(newVoucherUpdateBlock uint64, id int64) error {
	statement := squirrel.Update(blocksTable).
		Set(blocksVoucherUpdateBlockColumn, newVoucherUpdateBlock).
		Where(squirrel.Eq{blocksIdColumn: id})
	return q.db.Exec(statement)
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
