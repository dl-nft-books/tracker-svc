package token_listeners

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/token"
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

	defer func(iterator *token.TokencontractTransferIterator) {
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
	// Since l.to - l.from might exceed the max depth allowed in the chain,
	// we split the reading operation into several parallel processes
	// that are all sending caught events to the events channel

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

	var (
		wg        = new(sync.WaitGroup)
		intervals = helpers.SplitIntoIntervals(*l.from, *l.to, *l.maxDepth)
	)

	// Waitgroup is not necessary here, as we can simply run both listener and readers separately,
	// yet to just give a better sense of control over the parallel processing we will keep it as it is
	wg.Add(len(intervals))

	for _, interval := range intervals {
		go func(readerInterval helpers.Interval) {
			defer wg.Done()

			if tempErr := l.readTransfersInterval(readerInterval, ch); tempErr != nil {
				err = tempErr
				return
			}
		}(interval)
	}

	wg.Wait()
	return err
}

func (l *tokenListener) listenTransferEvents(ch chan<- etherdata.TransferEvent) (err error) {
	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := l.getWSInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *token.TokencontractTransfer)
	subscription, err := filterer.WatchTransfer(&opts, eventsChannel, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to watch transfer events")
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
