package data

type ChainStatistics struct {
	Id      int64 `db:"id" structs:"-" json:"-"`
	Amount  int64 `db:"amount" structs:"amount"`
	BookId  int64 `db:"book_id" structs:"book_id"`
	ChainId int64 `db:"chain_id" structs:"chain_id"`
}

type ChainStatisticsQ interface {
	New() ChainStatisticsQ

	FilterById(id ...int64) ChainStatisticsQ
	FilterByBookId(bookId ...int64) ChainStatisticsQ
	FilterByChainId(chainId ...int64) ChainStatisticsQ

	Get() (*ChainStatistics, error)
	Select() ([]ChainStatistics, error)
	Update(amount int64, id int64) error
	Insert(stats ChainStatistics) (id int64, err error)
	Delete(id int64) error
}
