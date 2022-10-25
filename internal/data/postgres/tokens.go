package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
)

const (
	tokensTable = "tokens"

	tokensId        = "id"
	tokensContract  = "contract"
	tokensName      = "name"
	tokensSymbol    = "symbol"
	tokensLastBlock = "last_block"
)

type tokensQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewTokensQ(database *pgdb.DB) data.TokensQ {
	return &tokensQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", tokensTable)).From(tokensTable),
	}
}

func (q *tokensQ) New() data.TokensQ {
	return NewTokensQ(q.database.Clone())
}

func (q *tokensQ) Page(page pgdb.OffsetPageParams) data.TokensQ {
	q.selector = page.ApplyTo(q.selector)
	return q
}

func (q *tokensQ) Get(id int64) (*data.Token, error) {
	var token data.Token

	err := q.database.Get(&token, q.selector.Where(squirrel.Eq{tokensId: id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &token, err
}

func (q *tokensQ) Insert(token data.Token) (id int64, err error) {
	statement := squirrel.Insert(tokensTable).
		Suffix("returning id").
		SetMap(structs.Map(&token))

	err = q.database.Get(&id, statement)
	return
}

func (q *tokensQ) UpdateLastBlock(lastBlock uint64, id int64) error {
	statement := squirrel.Update(tokensTable).
		Set(tokensLastBlock, lastBlock).
		Where(squirrel.Eq{tokensId: id})
	return q.database.Exec(statement)
}
