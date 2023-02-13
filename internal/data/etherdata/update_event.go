package etherdata

type UpdateEvent struct {
	Name        string
	Symbol      string
	Price       string
	FloorPrice  string
	BlockNumber uint64
}
