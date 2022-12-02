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
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/token"
)

var (
	NullIteratorErr = errors.New("iterator has a nil value")
)

type tokenIterator struct {
	rpc     *ethclient.Client
	from    *uint64
	to      *uint64
	address *common.Address
	ctx     context.Context

	processor ethereum.TokenProcessor
	converter converters.Converter

	// contractInstancesCache is a cache storing already initialized instances of contracts
	contractInstancesCache map[common.Address]*token.Tokencontract
}

func NewTokenContractReader(cfg config.Config, ctx context.Context) ethereum.TokenIterator {
	rpc := cfg.EtherClient().Rpc

	return &tokenIterator{
		rpc:                    rpc,
		ctx:                    ctx,
		contractInstancesCache: make(map[common.Address]*token.Tokencontract),

		converter: converters.NewConverter(rpc, ctx, cfg.NativeToken()),
	}
}

func (r *tokenIterator) From(from uint64) ethereum.TokenIterator {
	r.from = &from
	return r
}

func (r *tokenIterator) To(to uint64) ethereum.TokenIterator {
	r.to = &to
	return r
}

func (r *tokenIterator) WithAddress(address common.Address) ethereum.TokenIterator {
	r.address = &address
	return r
}

func (r *tokenIterator) WithCtx(ctx context.Context) ethereum.TokenIterator {
	r.ctx = ctx
	return r
}

func (r *tokenIterator) WithProcessor(processor ethereum.TokenProcessor) ethereum.TokenIterator {
	r.processor = processor
	return r
}

func (r *tokenIterator) validateParameters() error {
	if r.from == nil {
		return ethereum.FromNotSpecifiedErr
	}
	if r.address == nil {
		return ethereum.AddressNotSpecifiedErr
	}
	if r.processor == nil {
		return ethereum.ProcessorNotSpecifiedErr
	}

	return nil
}

func (r *tokenIterator) ProcessSuccessfulMints() (err error) {
	if err = r.validateParameters(); err != nil {
		return err
	}

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterSuccessfullyMinted(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		}, nil, nil,
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
			"depth":    *r.to - *r.from,
		})
	}

	defer func(iterator *token.TokencontractSuccessfullyMintedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event

		if raw != nil {
			var event *etherdata.SuccessfulMintEvent
			event, err = r.converter.SuccessfulMint(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert raw event format to a normal one")
			}

			if err = r.processor.ProcessSuccessfulMint(*event); err != nil {
				return errors.Wrap(err, "failed to process successful event")
			}
		}
	}

	return nil
}

func (r *tokenIterator) ProcessTransfers() (err error) {
	if err = r.validateParameters(); err != nil {
		return err
	}

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return errors.Wrap(err, "failed to create token contract instance")
	}

	iterator, err := instance.FilterTransfer(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		}, nil, nil, nil,
	)
	if iterator == nil {
		return errors.From(NullIteratorErr, logan.F{
			"contract": r.address.String(),
		})
	}

	defer func(iterator *token.TokencontractTransferIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			event := r.converter.Transfer(*raw)
			if err = r.processor.ProcessTransfer(event); err != nil {
				return errors.Wrap(err, "failed to process transfer event")
			}
		}
	}

	return nil
}

func (r *tokenIterator) ProcessUpdates() (err error) {
	if err = r.validateParameters(); err != nil {
		return err
	}

	instance, err := r.getContractInstance(*r.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a contract instance")
	}

	iterator, err := instance.FilterTokenContractParamsUpdated(
		&bind.FilterOpts{
			Start: *r.from,
			End:   r.to,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return NullIteratorErr
	}

	defer func(iterator *token.TokencontractTokenContractParamsUpdatedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			event := r.converter.Update(*raw)
			if err = r.processor.ProcessUpdate(event); err != nil {
				return errors.Wrap(err, "failed to process transfer event")
			}
		}
	}

	return nil
}

func (r *tokenIterator) getContractInstance(address common.Address) (*token.Tokencontract, error) {
	// Searching contract instance in cache, if not found -- create new and store
	cacheInstance, ok := r.contractInstancesCache[address]
	if ok {
		return cacheInstance, nil
	}

	newInstance, err := token.NewTokencontract(*r.address, r.rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize token factory instance for given address", logan.F{
			"address": address,
		})
	}

	r.contractInstancesCache[address] = newInstance
	return newInstance, nil
}
