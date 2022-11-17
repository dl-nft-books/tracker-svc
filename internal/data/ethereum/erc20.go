package ethereum

import "github.com/ethereum/go-ethereum/common"

var NullAddress = common.Address{}

// Erc20Info is a struct containing all the information
// retrieved from erc20 contract
type Erc20Info struct {
	TokenAddress common.Address `fig:"-"`
	Name         string         `fig:"name"`
	Symbol       string         `fig:"symbol"`
	Decimals     uint8          `fig:"decimals"`
}
