package factory_listeners

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/converters"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/factory"
	"sync"
)

type factoryListener struct {
	logger    *logan.Entry
	webSocket *ethclient.Client
	rpc       *ethclient.Client

	from     *uint64
	to       *uint64
	maxDepth *uint64

	address   *common.Address
	ctx       context.Context
	converter converters.EventConverter

	// rpcInstancesCache is a map storing already initialized rpc instances of contracts
	rpcInstancesCache map[common.Address]*factory.Tokenfactory
	rpcSync           sync.RWMutex

	// wsInstancesCache is a map storing already initialized ws instances of contracts
	wsInstancesCache map[common.Address]*factory.TokenfactoryFilterer
	wsSync           sync.RWMutex
}

func NewFactoryListener(cfg config.Config, ctx context.Context) ethereum.FactoryListener {
	rpc := cfg.EtherClient().Rpc

	return &factoryListener{
		logger:    cfg.Log(),
		webSocket: cfg.EtherClient().WebSocket,
		rpc:       rpc,
		ctx:       context.Background(),

		rpcInstancesCache: make(map[common.Address]*factory.Tokenfactory),
		rpcSync:           sync.RWMutex{},

		wsInstancesCache: map[common.Address]*factory.TokenfactoryFilterer{},
		wsSync:           sync.RWMutex{},

		converter: converters.NewEventConverter(rpc, ctx, cfg.NativeToken()),
	}
}

func (l *factoryListener) From(from uint64) ethereum.FactoryListener {
	l.from = &from
	return l
}

func (l *factoryListener) To(to uint64) ethereum.FactoryListener {
	l.to = &to
	return l
}

func (l *factoryListener) WithMaxDepth(maxDepth uint64) ethereum.FactoryListener {
	l.maxDepth = &maxDepth
	return l
}

func (l *factoryListener) WithAddress(address common.Address) ethereum.FactoryListener {
	l.address = &address
	return l
}

func (l *factoryListener) WithCtx(ctx context.Context) ethereum.FactoryListener {
	l.ctx = ctx
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

func (l *factoryListener) getRPCInstance(address common.Address) (*factory.Tokenfactory, error) {
	l.rpcSync.RLock()
	cacheInstance, ok := l.rpcInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := factory.NewTokenfactory(address, l.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize rpc token factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.rpcInstancesCache[address] = newInstance
	l.rpcSync.RUnlock()

	return newInstance, nil
}

func (l *factoryListener) getWSInstance(address common.Address) (*factory.TokenfactoryFilterer, error) {
	l.wsSync.RLock()
	cacheInstance, ok := l.wsInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := factory.NewTokenfactoryFilterer(address, l.webSocket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize ws token factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.wsInstancesCache[address] = newInstance
	l.wsSync.RUnlock()

	return newInstance, nil
}
