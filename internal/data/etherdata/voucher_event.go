package etherdata

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type VoucherUpdateEvent struct {
	VoucherTokenAddress common.Address
	VoucherTokenAmount  *big.Int
	BlockNumber         uint64
}
