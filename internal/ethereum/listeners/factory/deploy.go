package factory_listeners

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/factory"
)

func (l *factoryListener) readContractDeployedInterval(interval helpers.Interval, ch chan<- etherdata.ContractDeployedEvent) error {
	instance, err := l.getRPCInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterTokenContractDeployed(
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

	defer func(iterator *factory.TokenfactoryTokenContractDeployedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			var event *etherdata.ContractDeployedEvent
			event, err = l.converter.Deploy(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert raw event to the needed format")
			}

			ch <- *event
		}
	}

	return nil
}

func (l *factoryListener) readContractDeployedEvents(ch chan<- etherdata.ContractDeployedEvent) (err error) {
	// Since l.to - l.from might exceed the max depth allowed in the chain,
	// we split the reading operation into several smaller parallel processes
	// that are all sending caught events to the events channel

	lastChainBlock, err := l.rpc.BlockNumber(l.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block in chain")
	}

	if l.maxDepth == nil {
		return l.readContractDeployedInterval(helpers.Interval{
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

			if tempErr := l.readContractDeployedInterval(readerInterval, ch); tempErr != nil {
				err = tempErr
				return
			}
		}(interval)
	}

	wg.Wait()
	return err
}

func (l *factoryListener) listenContractCreatedEvents(ch chan<- etherdata.ContractDeployedEvent) (err error) {
	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := l.getWSInstance(*l.address)
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

			ch <- *convertedEvent
		}
	}
}

func (l *factoryListener) WatchContractCreatedEvents(ch chan<- etherdata.ContractDeployedEvent) (err error) {
	// We firstly range from l.from to the last chain block and then start a listener

	if err = l.validateParameters(); err != nil {
		return err
	}
	if err = l.readContractDeployedEvents(ch); err != nil {
		return errors.Wrap(err, "failed to read previous events")
	}
	if err = l.listenContractCreatedEvents(ch); err != nil {
		return errors.Wrap(err, "failed to listen to deploy events")
	}

	return nil
}
