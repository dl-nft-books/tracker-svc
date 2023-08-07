package data

type BookStatistics struct {
	Id       int64   `db:"id" structs:"-" json:"-"`
	Amount   int64   `db:"amount" structs:"amount"`
	UsdPrice float64 `db:"usd_price" structs:"usd_price"`
	BookId   int64   `db:"book_id" structs:"book_id"`
}

type BookStatisticsQ interface {
	New() BookStatisticsQ

	FilterById(id ...int64) BookStatisticsQ
	FilterByBookId(bookId ...int64) BookStatisticsQ

	Get() (*BookStatistics, error)
	Select() ([]BookStatistics, error)

	Insert(payment BookStatistics) (id int64, err error)
	Update(updateStatements BookStatistics, id int64) error
	Delete(id int64) error
}