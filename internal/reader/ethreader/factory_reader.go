package ethreader

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/reader"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokenfactory"
)

type FactoryContractReader struct {
	rpc *ethclient.Client

	from    *uint64
	to      *uint64
	address *common.Address
}

func NewFactoryContractReader(cfg config.Config) reader.FactoryReader {
	return &FactoryContractReader{
		rpc: cfg.EtherClient().Rpc,
	}
}

func (r *FactoryContractReader) From(from uint64) reader.FactoryReader {
	r.from = &from
	return r
}

func (r *FactoryContractReader) To(to uint64) reader.FactoryReader {
	r.to = &to
	return r
}

func (r *FactoryContractReader) WithAddress(address common.Address) reader.FactoryReader {
	r.address = &address
	return r
}

func (r *FactoryContractReader) validateParameters() error {
	if r.from == nil {
		return reader.FromNotSpecifiedErr
	}
	if r.address == nil {
		return reader.AddressNotSpecifiedErr
	}

	return nil
}

// GetContractCreatedEvents returns the deploy contract events
// in the form of ethereum.ContractCreatedEvent array
// based on contract, start and end blocks to search through
func (r *FactoryContractReader) GetContractCreatedEvents() ([]ethereum.ContractCreatedEvent, error) {
	events := make([]ethereum.ContractCreatedEvent, 0)

	if err := r.validateParameters(); err != nil {
		return nil, err
	}

	instance, err := tokenfactory.NewTokenfactory(*r.address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return nil, errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
		})
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		if event != nil {
			receipt, err := r.rpc.TransactionReceipt(context.Background(), event.Raw.TxHash)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			events = append(events, ethereum.ContractCreatedEvent{
				Address:     event.NewTokenContractAddr,
				BlockNumber: event.Raw.BlockNumber,
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
				Status:      receipt.Status,
			})
		}
	}

	return events, nil
}
