package etherdata

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// NFTRequestCreatedEvent is a struct containing all the data from
// the successful mint event which occurs after the user has
// successfully passed signature from the backend and minted a book
type NFTRequestCreatedEvent struct {
	RequestId     *big.Int
	Requester     common.Address
	TokenContract common.Address
	NftContract   common.Address
	NftId         *big.Int
	Status        uint64 // 0 if failed and 1 if successful
	BlockNumber   uint64
	Timestamp     time.Time
}
