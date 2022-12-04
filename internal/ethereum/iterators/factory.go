package iterators

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/converters"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/factory"
)

type factoryIterator struct {
	rpc     *ethclient.Client
	from    *uint64
	to      *uint64
	address *common.Address
	ctx     context.Context

	processor ethereum.FactoryProcessor
	converter converters.Converter

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*factory.Tokenfactory
}

func NewFactoryIterator(cfg config.Config, ctx context.Context) ethereum.FactoryIterator {
	rpc := cfg.EtherClient().Rpc

	return &factoryIterator{
		rpc:                    rpc,
		ctx:                    ctx,
		contractInstancesCache: make(map[common.Address]*factory.Tokenfactory),

		converter: converters.NewConverter(rpc, ctx, cfg.NativeToken()),
	}
}

func (r *factoryIterator) From(from uint64) ethereum.FactoryIterator {
	r.from = &from
	return r
}

func (r *factoryIterator) To(to uint64) ethereum.FactoryIterator {
	r.to = &to
	return r
}

func (r *factoryIterator) WithAddress(address common.Address) ethereum.FactoryIterator {
	r.address = &address
	return r
}

func (r *factoryIterator) WithCtx(ctx context.Context) ethereum.FactoryIterator {
	r.ctx = ctx
	return r
}

func (r *factoryIterator) WithProcessor(processor ethereum.FactoryProcessor) ethereum.FactoryIterator {
	r.processor = processor
	return r
}

func (r *factoryIterator) validateParameters() error {
	if r.from == nil {
		return ethereum.FromNotSpecifiedErr
	}
	if r.address == nil {
		return ethereum.AddressNotSpecifiedErr
	}
	if r.rpc == nil {
		return ethereum.RpcNotSpecifiedErr
	}

	return nil
}

func (r *factoryIterator) ProcessDeployEvents() (err error) {
	if err = r.validateParameters(); err != nil {
		return err
	}

	instance, err := r.getInstance(*r.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
		})
	}

	defer func(iterator *factory.TokenfactoryTokenContractDeployedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			var event *etherdata.ContractDeployedEvent
			event, err = r.converter.Deploy(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert raw event to the needed format")
			}

			if err = r.processor.ProcessDeploy(*event); err != nil {
				return errors.Wrap(err, "failed to process deploy event")
			}
		}
	}

	return nil
}

func (r *factoryIterator) getInstance(address common.Address) (*factory.Tokenfactory, error) {
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := factory.NewTokenfactory(address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
}
