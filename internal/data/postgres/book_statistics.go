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
	booksTable = "book_statistics"

	booksId       = "id"
	booksBookId   = "book_id"
	booksAmount   = "amount"
	booksUsdPrice = "usd_price"
)

type booksQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewBookStatisticsQ(database *pgdb.DB) data.BookStatisticsQ {
	return &booksQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", booksTable)).From(booksTable),
	}
}

func (q *booksQ) New() data.BookStatisticsQ {
	return NewBookStatisticsQ(q.database.Clone())
}

func (q *booksQ) Page(page pgdb.OffsetPageParams, sortByCols ...string) data.BookStatisticsQ {
	if len(sortByCols) == 0 {
		sortByCols = append(sortByCols, "id")
	}
	q.selector = page.ApplyTo(q.selector, sortByCols...)
	return q
}

func (q *booksQ) Sort(sort pgdb.Sorts) data.BookStatisticsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *booksQ) FilterById(id ...int64) data.BookStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{booksId: id})
	return q
}

func (q *booksQ) FilterByBookId(bookId ...int64) data.BookStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{booksBookId: bookId})
	return q
}

func (q *booksQ) Select() (books []data.BookStatistics, err error) {
	err = q.database.Select(&books, q.selector)
	return
}

func (q *booksQ) SelectWithChainId() (books []data.BookStatisticsWithChains, err error) {
	query := squirrel.Select(fmt.Sprintf("%s.*, min(chain_id) as chain_id", booksTable)).From(booksTable).
		Join(fmt.Sprintf("%s ON %s.book_id = %s.book_id", chainsTable, booksTable, chainsTable)).
		GroupBy(fmt.Sprintf("%s.id", booksTable))
	err = q.database.Select(&books, query)
	return
}

func (q *booksQ) Get() (*data.BookStatistics, error) {
	var book data.BookStatistics
	err := q.database.Get(&book, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &book, err
}

func (q *booksQ) Insert(book data.BookStatistics) (id int64, err error) {
	statement := squirrel.Insert(booksTable).
		Suffix("returning id").
		SetMap(structs.Map(&book))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *booksQ) Update(updateStatements data.BookStatistics, id int64) error {
	return q.database.Exec(squirrel.Update(booksTable).
		SetMap(map[string]interface{}{
			booksAmount:   updateStatements.Amount,
			booksUsdPrice: updateStatements.UsdPrice,
		}).
		Where(squirrel.Eq{booksId: id}))
}

func (q *booksQ) Delete(id int64) error {
	statement := squirrel.Delete(booksTable).Where(squirrel.Eq{booksId: id})
	return q.database.Exec(statement)
}
