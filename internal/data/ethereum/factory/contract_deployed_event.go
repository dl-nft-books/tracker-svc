package factory

import "github.com/ethereum/go-ethereum/common"

// ContractDeployedEvent is a struct containing all the data from
// contract deploy event
type ContractDeployedEvent struct {
	Address         common.Address // Address of a newly deployed contract
	BlockNumber     uint64
	Name, Symbol    string
	TokenId, Status uint64 // 0 if failed, 1 if successful
}
