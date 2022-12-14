package token_listeners

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/token"
	"time"
)

func (l *tokenListener) readSuccessfulMintInterval(interval helpers.Interval, ch chan<- etherdata.SuccessfulMintEvent) error {
	instance, err := l.getRPCInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterSuccessfullyMinted(
		&bind.FilterOpts{
			Start: interval.From,
			End:   &interval.To,
		}, nil, nil,
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

	defer func(iterator *token.TokencontractSuccessfullyMintedIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			var event *etherdata.SuccessfulMintEvent
			event, err = l.converter.SuccessfulMint(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert raw event to the needed format")
			}

			ch <- *event
		}
	}

	return nil
}

func (l *tokenListener) readSuccessfulMintEvents(ch chan<- etherdata.SuccessfulMintEvent) (err error) {
	lastChainBlock, err := l.rpc.BlockNumber(l.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block in chain")
	}

	if l.maxDepth == nil {
		return l.readSuccessfulMintInterval(helpers.Interval{
			From: *l.from,
			To:   lastChainBlock,
		}, ch)
	}

	if l.to == nil {
		l.to = &lastChainBlock
	}

	for _, interval := range helpers.SplitIntoIntervals(*l.from, *l.to, *l.maxDepth) {
		if err = l.readSuccessfulMintInterval(interval, ch); err != nil {
			return errors.Wrap(err, "failed to read mint interval", logan.F{
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

func (l *tokenListener) listenSuccessfulMintEvents(ch chan<- etherdata.SuccessfulMintEvent) (err error) {
	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := l.getWSInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *token.TokencontractSuccessfullyMinted)
	subscription, err := filterer.WatchSuccessfullyMinted(&opts, eventsChannel, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to watch succeesfully minted events")
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

			ch <- *convertedEvent
		}
	}
}

func (l *tokenListener) WatchSuccessfulMintEvents(ch chan<- etherdata.SuccessfulMintEvent) (err error) {
	// We firstly range from l.from to the last chain block and then start a listener

	if err = l.validateParameters(); err != nil {
		return err
	}
	if err = l.readSuccessfulMintEvents(ch); err != nil {
		return errors.Wrap(err, "failed to read previous events")
	}
	if err = l.listenSuccessfulMintEvents(ch); err != nil {
		return errors.Wrap(err, "failed to listen to deploy events")
	}

	return nil
}
