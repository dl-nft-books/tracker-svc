package data

//go:generate mockery --case=underscore --name=DB
type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Payments() PaymentsQ
	NftPayments() NftPaymentsQ
	Blocks() BlocksQ

	Transaction(func() error) error
}
