package etherdata

import (
	"math/big"
	"time"
)

// NFTRequestCanceledEvent is a struct containing all the data from
// the successful mint event which occurs after the user has
// successfully passed signature from the backend and minted a book
type NFTRequestCanceledEvent struct {
	RequestId   *big.Int
	Status      uint64 // 0 if failed and 1 if successful
	BlockNumber uint64
	Timestamp   time.Time
}
