package token

import "github.com/ethereum/go-ethereum/common"

// TransferEvent contains all the data from erc-721 transfer event
type TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId uint64
}
