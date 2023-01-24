package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	keyValueTable = "key_value"

	keyColumn     = "key"
	valueColumn   = "value"
	chainIdColumn = "chain_id"
)

var keyValueSelect = sq.Select("*").From(keyValueTable)

type keyValueQ struct {
	db *pgdb.DB
}

func NewKeyValueQ(db *pgdb.DB) data.KeyValueQ {
	return &keyValueQ{
		db: db,
	}
}

func (q *keyValueQ) Upsert(kv data.KeyValue) error {
	query := sq.Insert(keyValueTable).
		SetMap(structs.Map(kv)).
		Suffix("ON CONFLICT (key, chain_id) DO UPDATE SET value = EXCLUDED.value")

	return q.db.Exec(query)
}

func (q *keyValueQ) New() data.KeyValueQ {
	return NewKeyValueQ(q.db.Clone())
}

func (q *keyValueQ) Get(key string, chainId int64) (*data.KeyValue, error) {
	return q.get(key, chainId, false)
}

func (q *keyValueQ) LockingGet(key string, chainId int64) (*data.KeyValue, error) {
	return q.get(key, chainId, true)
}
func (q *keyValueQ) get(key string, chainId int64, forUpdate bool) (*data.KeyValue, error) {
	stmt := keyValueSelect.Where(sq.Eq{keyColumn: key, chainIdColumn: chainId})
	if forUpdate {
		stmt = stmt.Suffix("FOR UPDATE")
	}
	var value data.KeyValue

	err := q.db.Get(&value, stmt)
	fmt.Println(stmt.ToSql())
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &value, err
}
