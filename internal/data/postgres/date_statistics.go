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
	datesTable = "date_statistics"

	datesId     = "id"
	datesBookId = "book_id"
	datesAmount = "amount"
	datesDate   = "date"
)

type datesQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewDateStatisticsQ(database *pgdb.DB) data.DateStatisticsQ {
	return &datesQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", datesTable)).From(datesTable),
	}
}

func (q *datesQ) New() data.DateStatisticsQ {
	return NewDateStatisticsQ(q.database.Clone())
}

func (q *datesQ) Page(page pgdb.OffsetPageParams, sortByCols ...string) data.DateStatisticsQ {
	if len(sortByCols) == 0 {
		sortByCols = append(sortByCols, "id")
	}
	q.selector = page.ApplyTo(q.selector, sortByCols...)
	return q
}

func (q *datesQ) Sort(sort pgdb.Sorts) data.DateStatisticsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *datesQ) FilterById(id ...int64) data.DateStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{datesId: id})
	return q
}

func (q *datesQ) FilterByBookId(bookId ...int64) data.DateStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{datesBookId: bookId})
	return q
}

func (q *datesQ) FilterByDate(date ...string) data.DateStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{datesBookId: date})
	return q
}

func (q *datesQ) Select() (dates []data.DateStatistics, err error) {
	err = q.database.Select(&dates, q.selector)
	return
}

func (q *datesQ) Get() (*data.DateStatistics, error) {
	var date data.DateStatistics
	err := q.database.Get(&date, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &date, err
}

func (q *datesQ) Insert(date data.DateStatistics) (id int64, err error) {
	statement := squirrel.Insert(datesTable).
		Suffix("returning id").
		SetMap(structs.Map(&date))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *datesQ) Update(updateStatements data.DateStatistics, id int64) error {
	return q.database.Exec(squirrel.Update(datesDate).
		SetMap(structs.Map(&updateStatements)).
		Where(squirrel.Eq{datesId: id}))
}

func (q *datesQ) Delete(id int64) error {
	statement := squirrel.Delete(datesTable).Where(squirrel.Eq{datesId: id})
	return q.database.Exec(statement)
}
