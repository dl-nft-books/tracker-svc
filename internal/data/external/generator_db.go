package external

type GeneratorDB interface {
	New() GeneratorDB

	Tasks() TasksQ
	Tokens() TokensQ

	Transaction(func() error) error
}
