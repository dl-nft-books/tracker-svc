package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

const (
	tokensTable = "tokens"

	tokensId           = "id"
	tokensAccount      = "account"
	tokensTokenId      = "token_id"
	tokensBookId       = "book_id"
	tokensPaymentId    = "payment_id"
	tokensMetadataHash = "metadata_hash"
	tokensSignature    = "signature"
	tokensStatus       = "status"
)

type tokensQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewTokensQ(database *pgdb.DB) external.TokensQ {
	return &tokensQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tokensTable)).From(tokensTable),
	}
}

func (q *tokensQ) New() external.TokensQ {
	return NewTokensQ(q.database.Clone())
}

func (q *tokensQ) Page(page pgdb.OffsetPageParams) external.TokensQ {
	q.selector = page.ApplyTo(q.selector, "id")
	return q
}

func (q *tokensQ) Sort(sort pgdb.Sorts) external.TokensQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *tokensQ) FilterById(id ...int64) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensId: id})
	return q
}

func (q *tokensQ) FilterByStatus(status ...resources.TokenStatus) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensStatus: status})
	return q
}

func (q *tokensQ) FilterByAccount(account ...string) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensAccount: account})
	return q
}

func (q *tokensQ) FilterByTokenId(tokenId ...int64) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensTokenId: tokenId})
	return q
}

func (q *tokensQ) FilterByBookId(bookId ...int64) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensBookId: bookId})
	return q
}

func (q *tokensQ) FilterByPaymentId(paymentId ...int64) external.TokensQ {
	q.selector = q.selector.Where(squirrel.Eq{tokensPaymentId: paymentId})
	return q
}

func (q *tokensQ) Select() (tokens []external.Token, err error) {
	err = q.database.Select(&tokens, q.selector)
	return
}

func (q *tokensQ) Get() (*external.Token, error) {
	var token external.Token
	err := q.database.Get(&token, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &token, err
}

func (q *tokensQ) Insert(token external.Token) (id int64, err error) {
	statement := squirrel.Insert(tokensTable).
		Suffix("returning id").
		SetMap(structs.Map(&token))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *tokensQ) Transaction(fn func(q external.TokensQ) error) (err error) {
	return q.database.Transaction(func() error {
		return fn(q)
	})
}

func (q *tokensQ) UpdateStatus(newStatus resources.TokenStatus, id int64) error {
	statement := squirrel.Update(tokensTable).
		Set(tokensStatus, newStatus).
		Where(squirrel.Eq{tokensId: id})
	return q.database.Exec(statement)
}

func (q *tokensQ) UpdateAccount(account string, id int64) error {
	statement := squirrel.Update(tokensTable).
		Set(tokensAccount, account).
		Where(squirrel.Eq{tokensId: id})
	return q.database.Exec(statement)
}
