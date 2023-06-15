package resources

type NftRequestStatus int64

const (
	NftRequestPending NftRequestStatus = iota + 1
	NftRequestCanceled
	NftRequestAccepted
)

var nftRequestStates = map[NftRequestStatus]string{
	NftRequestPending:  "pending",
	NftRequestCanceled: "canceled",
	NftRequestAccepted: "accepted",
}

func (s NftRequestStatus) String() string {
	return nftRequestStates[s]
}
