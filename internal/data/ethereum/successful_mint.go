package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type SuccessfulMintEvent struct {
	Recipient         common.Address
	TokenId           int64
	Uri               string
	Erc20Info         Erc20Info
	Amount            *big.Int
	PaymentTokenPrice *big.Int
	MintedTokenPrice  *big.Int
	Status            uint64
	BlockNumber       uint64
	Timestamp         time.Time
}
