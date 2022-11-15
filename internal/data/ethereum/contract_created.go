package ethereum

import "github.com/ethereum/go-ethereum/common"

type ContractCreatedEvent struct {
	Address      common.Address
	BlockNumber  uint64
	Name, Symbol string
	Status       uint64
}
