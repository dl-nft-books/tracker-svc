package data

import "gitlab.com/distributed_lab/kit/pgdb"

type TokenStatistics struct {
	Id          int64   `db:"id" structs:"-" json:"-"`
	TokenSymbol string  `db:"token_symbol" structs:"token_symbol"`
	UsdPrice    float64 `db:"usd_price" structs:"usd_price"`
	TokenPrice  float64 `db:"token_price" structs:"token_price"`
	BookId      int64   `db:"book_id" structs:"book_id"`
}

type TokenStatisticsQ interface {
	New() TokenStatisticsQ

	FilterById(id ...int64) TokenStatisticsQ
	FilterByBookId(bookId ...int64) TokenStatisticsQ
	FilterByTokenSymbol(symbol ...string) TokenStatisticsQ

	Page(page pgdb.OffsetPageParams, sortByCols ...string) TokenStatisticsQ
	Sort(sort pgdb.Sorts) TokenStatisticsQ

	Get() (*TokenStatistics, error)
	Select() ([]TokenStatistics, error)
	Update(updateStatements TokenStatistics, id int64) error
	Insert(payment TokenStatistics) (id int64, err error)
	Delete(id int64) error
}
