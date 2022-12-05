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

func (l *tokenListener) readUpdatesInterval(interval helpers.Interval, ch chan<- etherdata.UpdateEvent) error {
	instance, err := l.getInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterTokenContractParamsUpdated(
		&bind.FilterOpts{
			Start: interval.From,
			End:   &interval.To,
		},
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

	defer func(iterator *token.TokencontractTokenContractParamsUpdatedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			ch <- l.converter.Update(*raw)
		}
	}

	return nil
}

func (l *tokenListener) readUpdateEvents(ch chan<- etherdata.UpdateEvent) (err error) {
	// Since l.to - l.from might exceed the max depth allowed in the chain,
	// we split the reading operation into several parallel processes
	// that are all sending caught events to the events channel

	lastChainBlock, err := l.rpc.BlockNumber(l.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block in chain")
	}

	if l.maxDepth == nil {
		return l.readUpdatesInterval(helpers.Interval{
			From: *l.from,
			To:   lastChainBlock,
		}, ch)
	}

	if l.to == nil {
		l.to = &lastChainBlock
	}

	var (
		wg        sync.WaitGroup
		intervals = helpers.SplitIntoIntervals(*l.from, *l.to, *l.maxDepth)
	)

	// Waitgroup is not necessary here, as we can simply run both listener and readers separately,
	// yet to just give a better sense of control over the parallel processing we will keep it as it is
	wg.Add(len(intervals))

	for _, interval := range intervals {
		go func(readerInterval helpers.Interval) {
			defer wg.Done()

			if tempErr := l.readUpdatesInterval(readerInterval, ch); tempErr != nil {
				err = tempErr
				return
			}
		}(interval)
	}

	wg.Wait()
	return err
}

func (l *tokenListener) listenUpdateEvents(ch chan<- etherdata.UpdateEvent) (err error) {
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
		return errors.Wrap(err, "failed to watch update events")
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

			ch <- l.converter.Update(*raw)
		}
	}
}

func (l *tokenListener) WatchUpdateEvents(ch chan<- etherdata.UpdateEvent) (err error) {
	// We firstly range from l.from to the last chain block and then start a listener

	if err = l.validateParameters(); err != nil {
		return err
	}
	if err = l.readUpdateEvents(ch); err != nil {
		return errors.Wrap(err, "failed to read previous events")
	}
	if err = l.listenUpdateEvents(ch); err != nil {
		return errors.Wrap(err, "failed to listen to update events")
	}

	return nil
}
