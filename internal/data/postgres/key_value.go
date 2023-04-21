package postgres

import (
	"database/sql"
	"fmt"
	"github.com/dl-nft-books/tracker-svc/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	keyValueTable = "key_value"

	keyColumn   = "key"
	valueColumn = "value"
)

type keyValueQ struct {
	db       *pgdb.DB
	selector sq.SelectBuilder
}

func NewKeyValueQ(db *pgdb.DB) data.KeyValueQ {
	return &keyValueQ{
		db:       db,
		selector: sq.Select("*").From(keyValueTable),
	}
}

func (q *keyValueQ) Upsert(kv data.KeyValue) error {
	query := sq.Insert(keyValueTable).
		SetMap(structs.Map(kv)).
		Suffix("ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value")

	return q.db.Exec(query)
}

func (q *keyValueQ) UpdateStatistics(kvs ...data.KeyValue) error {
	query := sq.Insert(keyValueTable)
	for _, kv := range kvs {
		query = query.Values(kv.Key, kv.Value)
	}
	query = query.Suffix("ON CONFLICT (key) DO UPDATE SET value = CAST(key_value.value as float) + CAST(EXCLUDED.value as float)")
	return q.db.Exec(query)
}

func (q *keyValueQ) New() data.KeyValueQ {
	return NewKeyValueQ(q.db.Clone())
}

func (q *keyValueQ) Get(key string) (*data.KeyValue, error) {
	return q.get(key, false)
}

func (q *keyValueQ) Select(key, notKey []string) (keyValues []data.KeyValue, err error) {
	orConditions := make(sq.Or, len(key))
	orConditionsNot := make(sq.Or, len(notKey))

	for i, k := range key {
		orConditions[i] = sq.Like{keyColumn: k}
	}
	for i, k := range notKey {
		orConditionsNot[i] = sq.NotLike{keyColumn: k}
	}
	if len(orConditions) > 0 {
		q.selector = q.selector.Where(orConditions)
	}
	if len(orConditionsNot) > 0 {
		q.selector = q.selector.Where(orConditionsNot)
	}
	fmt.Println(q.selector.ToSql())
	err = q.db.Select(&keyValues, q.selector)
	return
}

func (q *keyValueQ) LockingGet(key string) (*data.KeyValue, error) {
	return q.get(key, true)
}
func (q *keyValueQ) get(key string, forUpdate bool) (*data.KeyValue, error) {
	stmt := q.selector.Where(sq.Eq{keyColumn: key})
	if forUpdate {
		stmt = stmt.Suffix("FOR UPDATE")
	}
	var value data.KeyValue

	err := q.db.Get(&value, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &value, err
}
