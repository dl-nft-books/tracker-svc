package ethereum

import "github.com/ethereum/go-ethereum/common"

// ContractCreatedEvent is a struct containing all the data from
// contract deploy event
type ContractCreatedEvent struct {
	Address      common.Address // Address of a newly deployed contract
	BlockNumber  uint64
	Name, Symbol string
	Status       uint64 // 0 if failed, 1 if successful
}
