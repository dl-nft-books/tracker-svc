package eth_reader

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokenfactory"
)

type FactoryContractReader struct {
	rpc *ethclient.Client
}

func NewFactoryContractReader(rpc *ethclient.Client) FactoryContractReader {
	return FactoryContractReader{
		rpc: rpc,
	}
}

type ContractCreatedEvent struct {
	Address      common.Address
	BlockNumber  uint64
	Name, Symbol string
}

func (r *FactoryContractReader) GetEvents(
	contract common.Address,
	startBlock,
	endBlock uint64,
) (
	events []ContractCreatedEvent,
	lastBlock uint64,
	err error,
) {
	instance, err := tokenfactory.NewTokenfactory(contract, r.rpc)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize token factory instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
		&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		},
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, 0, NullIteratorErr
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		if event != nil {
			events = append(events, ContractCreatedEvent{
				Address:     event.NewTokenContractAddr,
				BlockNumber: event.Raw.BlockNumber,
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
			})
		}

		lastBlock = event.Raw.BlockNumber
	}

	return
}
