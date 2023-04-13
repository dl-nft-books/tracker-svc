package resources

type TokenPurchasedEventType int64

const (
	Native TokenPurchasedEventType = iota
	Erc20
	NFT
	Voucher
)

var eventTypes = map[TokenPurchasedEventType]string{
	Native:  "native",
	Erc20:   "erc20",
	NFT:     "nft",
	Voucher: "voucher",
}

func (s TokenPurchasedEventType) String() string {
	return eventTypes[s]
}
