package token

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

// SuccessfulMintEvent is a struct containing all the data from
// the successful mint event which occurs after the user has
// successfully passed signature from the backend and minted a book
type SuccessfulMintEvent struct {
	Recipient         common.Address     // Who obtained a erc-721 token
	TokenId           int64              // Token id assigned to the erc-721 token
	Uri               string             // Hash of a metadata file from the event
	Erc20Info         ethereum.Erc20Info // Erc20Info from the address used for mint payment
	Amount            *big.Int           // Amount of token paid
	PaymentTokenPrice *big.Int           // Price of a single token in dollars
	MintedTokenPrice  *big.Int           // Price of a minted erc-721 token in dollars
	Status            uint64             // 0 if failed and 1 if successful
	BlockNumber       uint64
	Timestamp         time.Time
}
