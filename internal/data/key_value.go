package data

type KeyValue struct {
	Key     string `db:"key" structs:"key"`
	Value   string `db:"value" structs:"value"`
	ChainId int64  `db:"chain_id" structs:"chain_id"`
}

//go:generate mockery --case=underscore --name=KeyValueQ
type KeyValueQ interface {
	New() KeyValueQ

	Get(key string, chainId int64) (*KeyValue, error)
	// Upsert updates value if there is one, insert if no
	Upsert(KeyValue) error
	// LockingGet reads row and locks the row for reading and updating
	// until the end of the current transaction
	LockingGet(key string, chainId int64) (*KeyValue, error)
}
