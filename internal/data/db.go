package data

//go:generate mockery --case=underscore --name=DB
type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Contracts() ContractsQ
	Payments() PaymentsQ
	Blocks() BlocksQ

	Transaction(func() error) error
}
