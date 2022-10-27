package postgres

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
)

type db struct {
	raw *pgdb.DB
}

func NewTrackerDB(rawDB *pgdb.DB) data.TrackerDB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.TrackerDB {
	return NewTrackerDB(db.raw.Clone())
}

func (db *db) KeyValue() data.KeyValueQ {
	return NewKeyValueQ(db.raw)
}

func (db *db) Contracts() data.ContractsQ {
	return NewTokensQ(db.raw)
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}
