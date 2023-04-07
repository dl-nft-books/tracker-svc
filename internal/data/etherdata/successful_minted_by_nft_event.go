package etherdata

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type SuccessfullyMintedByNftEvent struct {
	Recipient        common.Address // Who obtained a erc-721 marketplace
	TokenId          int64          // Token id assigned to the erc-721 marketplace
	Uri              string         // Hash of a metadata file from the event
	NftFloorPrice    *big.Int       // Price of a single marketplace in dollars
	MintedTokenPrice *big.Int       // Price of a minted erc-721 marketplace in dollars
	NftAddress       common.Address // Who obtained a erc-721 marketplace
	NftId            int64          // Token id assigned to the erc-721 marketplace
	Status           uint64         // 0 if failed and 1 if successful
	BlockNumber      uint64
	Timestamp        time.Time
}
