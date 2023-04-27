package data

type DateStatistics struct {
	Id     int64  `db:"id" structs:"-" json:"-"`
	Amount int64  `db:"amount" structs:"amount"`
	Date   string `db:"usd_price" structs:"usd_price"`
	BookId int64  `db:"book_id" structs:"book_id"`
}

type DateStatisticsQ interface {
	New() DateStatisticsQ

	FilterById(id ...int64) DateStatisticsQ
	FilterByBookId(bookId ...int64) DateStatisticsQ
	FilterByDate(date ...string) DateStatisticsQ

	Get() (*DateStatistics, error)
	Select() ([]DateStatistics, error)
	Update(updateStatements DateStatistics, id int64) error
	Insert(payment DateStatistics) (id int64, err error)
	Delete(id int64) error
}
