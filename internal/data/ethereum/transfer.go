package ethereum

import "github.com/ethereum/go-ethereum/common"

type TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId uint64
}
