package postgres

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
)

type generatorDB struct {
	raw *pgdb.DB
}

func NewGeneratorDB(rawDB *pgdb.DB) data.GeneratorDB {
	return &generatorDB{
		raw: rawDB,
	}
}

func (db *generatorDB) New() data.GeneratorDB {
	return NewGeneratorDB(db.raw.Clone())
}

func (db *generatorDB) Tasks() data.TasksQ {
	return NewTasksQ(db.raw)
}

func (db *generatorDB) Tokens() data.TokensQ {
	return NewTokensQ(db.raw)
}

func (db *generatorDB) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}
