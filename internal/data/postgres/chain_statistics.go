package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	chainsTable = "chain_statistics"

	chainsId      = "id"
	chainsBookId  = "book_id"
	chainsChainId = "chain_id"
	chainsAmount  = "amount"
)

type chainsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewChainStatisticsQ(database *pgdb.DB) data.ChainStatisticsQ {
	return &chainsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", chainsTable)).From(chainsTable),
	}
}

func (q *chainsQ) New() data.ChainStatisticsQ {
	return NewChainStatisticsQ(q.database.Clone())
}

func (q *chainsQ) Page(page pgdb.OffsetPageParams, sortByCols ...string) data.ChainStatisticsQ {
	if len(sortByCols) == 0 {
		sortByCols = append(sortByCols, "id")
	}
	q.selector = page.ApplyTo(q.selector, sortByCols...)
	return q
}

func (q *chainsQ) Sort(sort pgdb.Sorts) data.ChainStatisticsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *chainsQ) FilterById(id ...int64) data.ChainStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{chainsId: id})
	return q
}

func (q *chainsQ) FilterByBookId(bookId ...int64) data.ChainStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{chainsBookId: bookId})
	return q
}

func (q *chainsQ) FilterByChainId(chainId ...int64) data.ChainStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{chainsChainId: chainId})
	return q
}

func (q *chainsQ) Select() (chains []data.ChainStatistics, err error) {
	err = q.database.Select(&chains, q.selector)
	return
}

func (q *chainsQ) Get() (*data.ChainStatistics, error) {
	var chain data.ChainStatistics
	err := q.database.Get(&chain, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &chain, err
}

func (q *chainsQ) Insert(chain data.ChainStatistics) (id int64, err error) {
	statement := squirrel.Insert(chainsTable).
		Suffix("returning id").
		SetMap(structs.Map(&chain))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *chainsQ) Update(amount int64, id int64) error {
	return q.database.Exec(squirrel.Update(chainsTable).
		SetMap(map[string]interface{}{
			chainsAmount: amount,
		}).
		Where(squirrel.Eq{chainsId: id}))
}

func (q *chainsQ) Delete(id int64) error {
	statement := squirrel.Delete(chainsTable).Where(squirrel.Eq{chainsId: id})
	return q.database.Exec(statement)
}
