package ethereum

import "github.com/ethereum/go-ethereum/common"

type Erc20Info struct {
	TokenAddress common.Address
	Name         string
	Symbol       string
	Decimals     uint8
}
