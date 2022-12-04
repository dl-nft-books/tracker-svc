package readers

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

type factoryListener struct {
	webSocket *ethclient.Client
	rpc       *ethclient.Client
	address   *common.Address
	ctx       context.Context

	processor ethereum.FactoryProcessor
	converter converters.Converter

	// contractInstancesCache is a map storing already initialized instances of contracts
	instancesCache map[common.Address]*factory.Tokenfactory
}

func NewFactoryListener(cfg config.Config, ctx context.Context) ethereum.FactoryListener {
	rpc := cfg.EtherClient().Rpc

	return &factoryListener{
		webSocket: cfg.EtherClient().WebSocket,
		rpc:       rpc,
		ctx:       context.Background(),

		instancesCache: make(map[common.Address]*factory.Tokenfactory),
		converter:      converters.NewConverter(rpc, ctx, cfg.NativeToken()),
	}
}

func (l *factoryListener) WithAddress(address common.Address) ethereum.FactoryListener {
	l.address = &address
	return l
}

func (l *factoryListener) WithCtx(ctx context.Context) ethereum.FactoryListener {
	l.ctx = ctx
	return l
}

func (l *factoryListener) WithProcessor(processor ethereum.FactoryProcessor) ethereum.FactoryListener {
	l.processor = processor
	return l
}

func (l *factoryListener) validateParameters() error {
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

func (l *factoryListener) WatchContractCreatedEvents() (err error) {
	if err = l.validateParameters(); err != nil {
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
			convertedEvent, err = l.converter.Deploy(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert event to the needed type")
			}

			if err = l.processor.ProcessDeploy(*convertedEvent); err != nil {
				return errors.Wrap(err, "failed to process an event")
			}
		}
	}
}

func (r *factoryListener) getInstance(address common.Address) (*factory.Tokenfactory, error) {
	cacheInstance, ok := r.instancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := factory.NewTokenfactory(address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.instancesCache[address] = newInstance
	return newInstance, nil
}
