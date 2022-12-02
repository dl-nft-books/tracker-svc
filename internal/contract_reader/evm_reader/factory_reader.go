package evm_based_reader

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/contract-reader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum/factory"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokenfactory"
)

type FactoryContractReader struct {
	rpc *ethclient.Client

	from    *uint64
	to      *uint64
	address *common.Address
	ctx     context.Context

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*tokenfactory.Tokenfactory
}

func NewFactoryContractReader(cfg config.Config) contract_reader.TokenFactoryReader {
	return &FactoryContractReader{
		rpc:                    cfg.EtherClient().Rpc,
		ctx:                    context.Background(),
		contractInstancesCache: make(map[common.Address]*tokenfactory.Tokenfactory),
	}
}

func (r *FactoryContractReader) From(from uint64) contract_reader.TokenFactoryReader {
	r.from = &from
	return r
}

func (r *FactoryContractReader) To(to uint64) contract_reader.TokenFactoryReader {
	r.to = &to
	return r
}

func (r *FactoryContractReader) WithAddress(address common.Address) contract_reader.TokenFactoryReader {
	r.address = &address
	return r
}

func (r *FactoryContractReader) WithCtx(ctx context.Context) contract_reader.TokenFactoryReader {
	r.ctx = ctx
	return r
}

func (r *FactoryContractReader) validateParameters() error {
	if r.from == nil {
		return contract_reader.FromNotSpecifiedErr
	}
	if r.address == nil {
		return contract_reader.AddressNotSpecifiedErr
	}

	return nil
}

// GetContractCreatedEvents returns the deploy contract events
// in the form of ethereum.ContractDeployedEvent array
// based on contract, start and end blocks to search through
func (r *FactoryContractReader) GetContractCreatedEvents() (events []factory.ContractDeployedEvent, err error) {
	if err = r.validateParameters(); err != nil {
		return nil, err
	}

	instance, err := r.getInstance(*r.address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
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

	defer func(iterator *tokenfactory.TokenfactoryTokenContractDeployedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		event := iterator.Event
		if event != nil {
			receipt, err := r.rpc.TransactionReceipt(r.ctx, event.Raw.TxHash)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get tx receipt", logan.F{
					"tx_hash": event.Raw.TxHash.String(),
				})
			}

			events = append(events, factory.ContractDeployedEvent{
				Address:     event.NewTokenContractAddr,
				BlockNumber: event.Raw.BlockNumber,
				Name:        event.TokenName,
				Symbol:      event.TokenSymbol,
				Status:      receipt.Status,
				TokenId:     event.TokenContractId.Uint64(),
			})
		}
	}

	return events, nil
}

func (r *FactoryContractReader) getInstance(address common.Address) (*tokenfactory.Tokenfactory, error) {
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := tokenfactory.NewTokenfactory(address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
}
