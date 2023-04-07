package token_listeners

import (
	"context"
	"github.com/dl-nft-books/tracker-svc/internal/data/etherdata"
	"github.com/dl-nft-books/network-svc/connector/models"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/tracker-svc/internal/ethereum"
	"github.com/dl-nft-books/tracker-svc/internal/ethereum/converters"
	"github.com/dl-nft-books/tracker-svc/solidity/generated/marketplace"
)

type tokenListener struct {
	webSocket *ethclient.Client
	rpc       *ethclient.Client
	logger    *logan.Entry

	from                  *uint64
	to                    *uint64
	maxDepth              *uint64
	delayBetweenIntervals *time.Duration

	address   *common.Address
	ctx       context.Context
	mutex     *sync.RWMutex
	converter converters.EventConverter

	// rpcInstancesCache is a map storing already initialized instances of contracts
	rpcInstancesCache map[common.Address]*marketplace.Marketplace
	// wsInstancesCache is a map storing already initialized ws instances of contracts
	wsInstancesCache map[common.Address]*marketplace.MarketplaceFilterer
}

func NewTokenListener(ctx context.Context, mutex *sync.RWMutex, network models.NetworkDetailedResponse) ethereum.TokenListener {
	rpc := network.RpcUrl
	nativeToken := etherdata.Erc20Info{
		Name:     network.TokenName,
		Symbol:   network.TokenSymbol,
		Decimals: uint8(network.Decimals),
	}
	return &tokenListener{
		webSocket: network.WsUrl,
		rpc:       rpc,
		ctx:       ctx,
		mutex:     mutex,

		rpcInstancesCache: make(map[common.Address]*marketplace.Marketplace),
		wsInstancesCache:  make(map[common.Address]*marketplace.MarketplaceFilterer),

		converter: converters.NewEventConverter(rpc, ctx, nativeToken),
	}
}

func (l *tokenListener) From(from uint64) ethereum.TokenListener {
	l.from = &from
	return l
}

func (l *tokenListener) To(to uint64) ethereum.TokenListener {
	l.to = &to
	return l
}

func (l *tokenListener) WithMaxDepth(maxDepth uint64) ethereum.TokenListener {
	l.maxDepth = &maxDepth
	return l
}

func (l *tokenListener) WithAddress(address common.Address) ethereum.TokenListener {
	l.address = &address
	return l
}

func (l *tokenListener) WithCtx(ctx context.Context) ethereum.TokenListener {
	l.ctx = ctx
	return l
}

func (l *tokenListener) WithDelayBetweenIntervals(delay time.Duration) ethereum.TokenListener {
	l.delayBetweenIntervals = &delay
	return l
}

func (l *tokenListener) validateParameters() error {
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

func (l *tokenListener) getRPCInstance(address common.Address) (*marketplace.Marketplace, error) {
	l.mutex.RLock()
	cacheInstance, ok := l.rpcInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := marketplace.NewMarketplace(address, l.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize rpc marketplace factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.rpcInstancesCache[address] = newInstance
	l.mutex.RUnlock()

	return newInstance, nil
}

func (l *tokenListener) getWSInstance(address common.Address) (*marketplace.MarketplaceFilterer, error) {
	l.mutex.RLock()
	cacheInstance, ok := l.wsInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := marketplace.NewMarketplaceFilterer(address, l.webSocket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize ws marketplace factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.wsInstancesCache[address] = newInstance
	l.mutex.RUnlock()

	return newInstance, nil
}
