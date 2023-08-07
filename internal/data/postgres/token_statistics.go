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
	tokensTable = "token_statistics"

	tokensId          = "id"
	tokensBookId      = "book_id"
	tokensTokenSymbol = "token_symbol"
	tokensUsdPrice    = "usd_price"
	tokensTokenPrice  = "token_price"
)

type tokensQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewTokenStatisticsQ(database *pgdb.DB) data.TokenStatisticsQ {
	return &tokensQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tokensTable)).From(tokensTable),
	}
}

func (q *tokensQ) New() data.TokenStatisticsQ {
	return NewTokenStatisticsQ(q.database.Clone())
}

func (q *tokensQ) Page(page pgdb.OffsetPageParams, sortByCols ...string) data.TokenStatisticsQ {
	if len(sortByCols) == 0 {
		sortByCols = append(sortByCols, "id")
	}
	q.selector = page.ApplyTo(q.selector, sortByCols...)
	return q
}

func (q *tokensQ) Sort(sort pgdb.Sorts) data.TokenStatisticsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *tokensQ) FilterById(id ...int64) data.TokenStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensId: id})
	return q
}

func (q *tokensQ) FilterByBookId(bookId ...int64) data.TokenStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensBookId: bookId})
	return q
}

func (q *tokensQ) FilterByTokenSymbol(symbol ...string) data.TokenStatisticsQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensTokenSymbol: symbol})
	return q
}

func (q *tokensQ) Select() (tokens []data.TokenStatistics, err error) {
	err = q.database.Select(&tokens, q.selector)
	return
}

func (q *tokensQ) Get() (*data.TokenStatistics, error) {
	var token data.TokenStatistics
	err := q.database.Get(&token, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &token, err
}

func (q *tokensQ) Insert(token data.TokenStatistics) (id int64, err error) {
	statement := squirrel.Insert(tokensTable).
		Suffix("returning id").
		SetMap(structs.Map(&token))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *tokensQ) Update(updateStatements data.TokenStatistics, id int64) error {
	return q.database.Exec(squirrel.Update(tokensTable).
		SetMap(map[string]interface{}{
			tokensTokenPrice: updateStatements.TokenPrice,
			tokensUsdPrice:   updateStatements.UsdPrice,
		}).
		Where(squirrel.Eq{tokensId: id}))
}

func (q *tokensQ) Delete(id int64) error {
	statement := squirrel.Delete(tokensTable).Where(squirrel.Eq{tokensId: id})
	return q.database.Exec(statement)
}
