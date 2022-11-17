package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"
)

const booksTableName = "book"
const priceColumnName = "price"

func NewBooksQ(db *pgdb.DB) external.BookQ {
	return &BooksQ{
		db: db.Clone(),
		sql: squirrel.
			Select("*").
			From(fmt.Sprintf("%s as b", booksTableName)),
	}
}

type BooksQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func (b *BooksQ) New() external.BookQ {
	return NewBooksQ(b.db)
}

func (b *BooksQ) Get() (*external.Book, error) {
	var result external.Book

	err := b.db.Get(&result, b.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *BooksQ) Select() (books []external.Book, err error) {
	err = q.db.Select(&books, q.sql)
	return
}

func (b *BooksQ) FilterByID(id ...int64) external.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.id": id})
	return b
}

func (b *BooksQ) FilterActual() external.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.deleted": "f"})
	return b
}

func (b *BooksQ) FilterByContractAddress(address ...string) external.BookQ {
	b.sql = b.sql.Where(squirrel.Eq{"b.contract_address": address})
	return b
}
