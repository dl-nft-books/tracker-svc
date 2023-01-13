package token_listeners

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/converters"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokencontract"
)

type tokenListener struct {
	webSocket *ethclient.Client
	rpc       *ethclient.Client

	from                  *uint64
	to                    *uint64
	maxDepth              *uint64
	delayBetweenIntervals *time.Duration

	address   *common.Address
	ctx       context.Context
	mutex     *sync.RWMutex
	converter converters.EventConverter

	// rpcInstancesCache is a map storing already initialized instances of contracts
	rpcInstancesCache map[common.Address]*tokencontract.Tokencontract
	// wsInstancesCache is a map storing already initialized ws instances of contracts
	wsInstancesCache map[common.Address]*tokencontract.TokencontractFilterer
}

func NewTokenListener(cfg config.Config, ctx context.Context, mutex *sync.RWMutex) ethereum.TokenListener {
	rpc := cfg.EtherClient().Rpc

	return &tokenListener{
		webSocket: cfg.EtherClient().WebSocket,
		rpc:       rpc,
		ctx:       ctx,
		mutex:     mutex,

		rpcInstancesCache: make(map[common.Address]*tokencontract.Tokencontract),
		wsInstancesCache:  make(map[common.Address]*tokencontract.TokencontractFilterer),

		converter: converters.NewEventConverter(rpc, ctx, cfg.NativeToken()),
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

func (l *tokenListener) getRPCInstance(address common.Address) (*tokencontract.Tokencontract, error) {
	l.mutex.RLock()
	cacheInstance, ok := l.rpcInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := tokencontract.NewTokencontract(address, l.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize rpc token factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.rpcInstancesCache[address] = newInstance
	l.mutex.RUnlock()

	return newInstance, nil
}

func (l *tokenListener) getWSInstance(address common.Address) (*tokencontract.TokencontractFilterer, error) {
	l.mutex.RLock()
	cacheInstance, ok := l.wsInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := tokencontract.NewTokencontractFilterer(address, l.webSocket)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize ws token factory instance for given address", logan.F{
			"address": address,
		})
	}

	l.wsInstancesCache[address] = newInstance
	l.mutex.RUnlock()

	return newInstance, nil
}
