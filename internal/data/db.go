package data

//go:generate mockery --case=underscore --name=DB

type StatisticsQ struct {
	BookStatisticsQ
	TokenStatisticsQ
	DateStatisticsQ
	ChainStatisticsQ
}

type DB interface {
	New() DB

	KeyValue() KeyValueQ
	Payments() PaymentsQ
	NftRequests() NftRequestsQ
	Blocks() BlocksQ
	Statistics() StatisticsQ

	Transaction(func() error) error
}
