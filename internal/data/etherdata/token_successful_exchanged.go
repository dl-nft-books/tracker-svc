package etherdata

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// TokenSuccessfullyExchangedEvent is a struct containing all the data from
// the successful mint event which occurs after the user has
// successfully passed signature from the backend and minted a book
type TokenSuccessfullyExchangedEvent struct {
	ContractAddress common.Address // address of book
	Recipient       common.Address // Who obtained a erc-721 marketplace
	TokenId         int64          // Token id assigned to the erc-721 marketplace
	NftId           int64          // Token id assigned to the erc-721 marketplace
	RequestId       int64          // Token id assigned to the erc-721 marketplace
	Uri             string         // Hash of a metadata file from the event
	Status          uint64         // 0 if failed and 1 if successful
	BlockNumber     uint64
	Timestamp       time.Time
}
