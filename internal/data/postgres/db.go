package postgres

import (
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type db struct {
	raw *pgdb.DB
}

func NewDB(rawDB *pgdb.DB) data.DB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.DB {
	return NewDB(db.raw.Clone())
}

func (db *db) KeyValue() data.KeyValueQ {
	return NewKeyValueQ(db.raw)
}

func (db *db) Payments() data.PaymentsQ {
	return NewPaymentsQ(db.raw)
}

func (db *db) Blocks() data.BlocksQ {
	return NewBlocksQ(db.raw)
}

func (db *db) Statistics() data.StatisticsQ {
	return data.StatisticsQ{
		BookStatisticsQ:  NewBookStatisticsQ(db.raw),
		TokenStatisticsQ: NewTokenStatisticsQ(db.raw),
		DateStatisticsQ:  NewDateStatisticsQ(db.raw),
		ChainStatisticsQ: NewChainStatisticsQ(db.raw),
	}
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}
