package token_listeners

import (
	"github.com/ethereum/go-ethereum/log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokencontract"
)

func (l *tokenListener) readTransfersInterval(interval helpers.Interval, ch chan<- etherdata.TransferEvent) error {
	instance, err := l.getRPCInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterTransfer(
		&bind.FilterOpts{
			Start: interval.From,
			End:   &interval.To,
		}, nil, nil, nil,
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize an iterator")
	}
	if iterator == nil {
		return errors.From(ethereum.NullIteratorErr, logan.F{
			"contract": l.address.String(),
			"depth":    interval.To - interval.From,
		})
	}

	defer func(iterator *tokencontract.TokencontractTransferIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			ch <- l.converter.Transfer(*raw)
		}
	}

	return nil
}

func (l *tokenListener) readTransferEvents(ch chan<- etherdata.TransferEvent) (err error) {
	lastChainBlock, err := l.rpc.BlockNumber(l.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block in chain")
	}

	if l.maxDepth == nil {
		return l.readTransfersInterval(helpers.Interval{
			From: *l.from,
			To:   lastChainBlock,
		}, ch)
	}

	if l.to == nil {
		l.to = &lastChainBlock
	}

	for _, interval := range helpers.SplitIntoIntervals(*l.from, *l.to, *l.maxDepth) {
		if err = l.readTransfersInterval(interval, ch); err != nil {
			return errors.Wrap(err, "failed to read interval", logan.F{
				"from": interval.From,
				"to":   interval.To,
			})
		}

		if l.delayBetweenIntervals != nil {
			time.Sleep(*l.delayBetweenIntervals)
		}
	}

	return nil
}

func (l *tokenListener) listenTransferEvents(ch chan<- etherdata.TransferEvent) (err error) {
	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := l.getWSInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *tokencontract.TokencontractTransfer)
	subscription, err := filterer.WatchTransfer(&opts, eventsChannel, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to watch transfer events")
	}

	for {
		select {
		case err = <-subscription.Err():
			//subscription.Unsubscribe()
			log.Error(errors.Wrap(err, "failed to listen to a subscription").Error())
			continue
		case raw := <-eventsChannel:
			if raw == nil {
				continue
			}

			ch <- l.converter.Transfer(*raw)
		}
	}
}

func (l *tokenListener) WatchTransferEvents(ch chan<- etherdata.TransferEvent) (err error) {
	// We firstly range from l.from to the last chain block and then start a listener

	if err = l.validateParameters(); err != nil {
		return err
	}
	if err = l.readTransferEvents(ch); err != nil {
		return errors.Wrap(err, "failed to read previous events")
	}
	if err = l.listenTransferEvents(ch); err != nil {
		return errors.Wrap(err, "failed to listen to transfer events")
	}

	return nil
}
