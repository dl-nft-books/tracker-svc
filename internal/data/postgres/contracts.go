package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
)

const (
	contractsTable = "contracts"

	contractsId        = "id"
	contractsContract  = "contract"
	contractsName      = "name"
	contractsSymbol    = "symbol"
	contractsLastBlock = "last_block"
)

type contractsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewContractsQ(database *pgdb.DB) data.ContractsQ {
	return &contractsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", contractsTable)).From(contractsTable),
	}
}

func (q *contractsQ) New() data.ContractsQ {
	return NewContractsQ(q.database.Clone())
}

func (q *contractsQ) Page(page pgdb.OffsetPageParams) data.ContractsQ {
	q.selector = page.ApplyTo(q.selector)
	return q
}

func (q *contractsQ) Get(id int64) (*data.Contract, error) {
	var token data.Contract

	err := q.database.Get(&token, q.selector.Where(squirrel.Eq{contractsId: id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &token, err
}

func (q *contractsQ) Insert(contracts ...data.Contract) (id []int64, err error) {
	query := squirrel.Insert(contractsTable).Columns(structs.Names(contracts)...)
	for _, contract := range contracts {
		query = query.Values(structs.Values(contract)...)
	}

	query = query.Suffix("returning id")
	err = q.database.Select(&id, query)
	return
}

func (q *contractsQ) UpdateLastBlock(lastBlock uint64, id int64) error {
	statement := squirrel.Update(contractsTable).
		Set(contractsLastBlock, lastBlock).
		Where(squirrel.Eq{contractsId: id})
	return q.database.Exec(statement)
}

func (q *contractsQ) Select() (contracts []data.Contract, err error) {
	err = q.database.Select(&contracts, q.selector)
	return
}
