package data

//go:generate mockery --case=underscore --name=DB
type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Contracts() ContractsQ

	Transaction(func() error) error
}
