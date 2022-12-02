package readers

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata/factory"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/converters"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/factory"
)

type factoryContractListener struct {
	webSocket *ethclient.Client
	rpc       *ethclient.Client
	address   *common.Address
	ctx       context.Context
	processor ethereum.TokenFactoryProcessor

	// contractInstancesCache is a map storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*factory.Tokenfactory
}

func NewFactoryContractListener(cfg config.Config) ethereum.TokenFactoryListener {
	return &factoryContractListener{
		webSocket: cfg.EtherClient().WebSocket,
		rpc:       cfg.EtherClient().Rpc,
		ctx:       context.Background(),

		contractInstancesCache: make(map[common.Address]*factory.Tokenfactory),
	}
}

func (l *factoryContractListener) WithAddress(address common.Address) ethereum.TokenFactoryListener {
	l.address = &address
	return l
}

func (l *factoryContractListener) WithCtx(ctx context.Context) ethereum.TokenFactoryListener {
	l.ctx = ctx
	return l
}

func (l *factoryContractListener) WithProcessor(processor ethereum.TokenFactoryProcessor) ethereum.TokenFactoryListener {
	l.processor = processor
	return l
}

func (l *factoryContractListener) validateParameters() error {
	if l.address == nil {
		return ethereum.AddressNotSpecifiedErr
	}
	if l.webSocket == nil {
		return ethereum.WSNotSpecifiedErr
	}
	if l.rpc == nil {
		return ethereum.RpcNotSpecifiedErr
	}

	return nil
}

func (l *factoryContractListener) WatchContractCreatedEvents() error {
	if err := l.validateParameters(); err != nil {
		return err
	}

	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := factory.NewTokenfactoryFilterer(*l.address, l.webSocket)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *factory.TokenfactoryTokenContractDeployed)

	subscription, err := filterer.WatchTokenContractDeployed(&opts, eventsChannel)
	if err != nil {
		return errors.Wrap(err, "failed to watch contracts deployed")
	}

	for {
		select {
		case err = <-subscription.Err():
			subscription.Unsubscribe()
			return errors.Wrap(err, "failed to listen to a subscription")
		case raw := <-eventsChannel:
			if raw == nil {
				continue
			}

			var convertedEvent *etherdata.ContractDeployedEvent
			convertedEvent, err = converters.ConvertDeployType(*raw, l.rpc, l.ctx)
			if err != nil {
				return errors.Wrap(err, "failed to convert event to the needed type")
			}

			if err = l.processor.ProcessDeploy(*convertedEvent); err != nil {
				return errors.Wrap(err, "failed to process an event")
			}
		}
	}
}

func (r *factoryContractListener) getInstance(address common.Address) (*factory.Tokenfactory, error) {
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
