package readers

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum/converters"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/token"
)

type tokenListener struct {
	webSocket *ethclient.Client
	rpc       *ethclient.Client
	address   *common.Address
	ctx       context.Context

	processor ethereum.TokenProcessor
	converter converters.Converter

	// instancesCache is a map storing already initialized instances of contracts
	instancesCache map[common.Address]*token.Tokencontract
}

func NewTokenListener(cfg config.Config, ctx context.Context) ethereum.TokenListener {
	rpc := cfg.EtherClient().Rpc

	return &tokenListener{
		webSocket: cfg.EtherClient().WebSocket,
		rpc:       rpc,
		ctx:       context.Background(),

		instancesCache: make(map[common.Address]*token.Tokencontract),
		converter:      converters.NewConverter(rpc, ctx, cfg.NativeToken()),
	}
}

func (l *tokenListener) WithAddress(address common.Address) ethereum.TokenListener {
	l.address = &address
	return l
}

func (l *tokenListener) WithCtx(ctx context.Context) ethereum.TokenListener {
	l.ctx = ctx
	return l
}

func (l *tokenListener) WithProcessor(processor ethereum.TokenProcessor) ethereum.TokenListener {
	l.processor = processor
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

func (l *tokenListener) WatchSuccessfulMintEvents() (err error) {
	if err = l.validateParameters(); err != nil {
		return err
	}

	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := token.NewTokencontractFilterer(*l.address, l.webSocket)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *token.TokencontractSuccessfullyMinted)

	subscription, err := filterer.WatchSuccessfullyMinted(&opts, eventsChannel, nil, nil)
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

			var convertedEvent *etherdata.SuccessfulMintEvent
			convertedEvent, err = l.converter.SuccessfulMint(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert event to the needed type")
			}

			if err = l.processor.ProcessSuccessfulMint(*convertedEvent); err != nil {
				return errors.Wrap(err, "failed to process an event")
			}
		}
	}
}

func (l *tokenListener) WatchTransferEvents() (err error) {
	if err = l.validateParameters(); err != nil {
		return err
	}

	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := token.NewTokencontractFilterer(*l.address, l.webSocket)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *token.TokencontractTransfer)

	subscription, err := filterer.WatchTransfer(&opts, eventsChannel, nil, nil, nil)
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

			event := l.converter.Transfer(*raw)
			if err = l.processor.ProcessTransfer(event); err != nil {
				return errors.Wrap(err, "failed to process an event")
			}
		}
	}
}

func (l *tokenListener) WatchUpdateEvents() (err error) {
	if err = l.validateParameters(); err != nil {
		return err
	}

	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := token.NewTokencontractFilterer(*l.address, l.webSocket)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *token.TokencontractTokenContractParamsUpdated)

	subscription, err := filterer.WatchTokenContractParamsUpdated(&opts, eventsChannel)
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

			event := l.converter.Update(*raw)
			if err = l.processor.ProcessUpdate(event); err != nil {
				return errors.Wrap(err, "failed to process an event")
			}
		}
	}
}
