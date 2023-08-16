package data

type BookStatistics struct {
	Id       int64  `db:"id" structs:"-" json:"-"`
	Amount   int64  `db:"amount" structs:"amount"`
	UsdPrice string `db:"usd_price" structs:"usd_price"`
	BookId   int64  `db:"book_id" structs:"book_id"`
}

type BookStatisticsWithChains struct {
	BookStatistics
	ChainId int64 `db:"chain_id" structs:"chain_id"`
}

type BookStatisticsQ interface {
	New() BookStatisticsQ

	FilterById(id ...int64) BookStatisticsQ
	FilterByBookId(bookId ...int64) BookStatisticsQ

	Get() (*BookStatistics, error)
	Select() ([]BookStatistics, error)
	SelectWithChainId() (books []BookStatisticsWithChains, err error)

	Insert(payment BookStatistics) (id int64, err error)
	Update(updateStatements BookStatistics, id int64) error
	Delete(id int64) error
}
