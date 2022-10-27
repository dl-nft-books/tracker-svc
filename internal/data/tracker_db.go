package data

//go:generate mockery --case=underscore --name=DB
type TrackerDB interface {
	New() TrackerDB

	KeyValue() KeyValueQ
	Contracts() ContractsQ

	Transaction(func() error) error
}
