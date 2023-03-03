package token_listeners

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/helpers"
	"gitlab.com/tokend/nft-books/contract-tracker/solidity/generated/tokencontract"
)

func (l *tokenListener) readSuccessfulMintByNftInterval(interval helpers.Interval, ch chan<- etherdata.SuccessfullyMintedByNftEvent) error {
	instance, err := l.getRPCInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to get instance")
	}

	iterator, err := instance.FilterSuccessfullyMintedByNFT(
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

	defer func(iterator *tokencontract.TokencontractSuccessfullyMintedByNFTIterator) {
		if tempErr := iterator.Close(); tempErr != nil {
			err = tempErr
		}
	}(iterator)

	for iterator.Next() {
		raw := iterator.Event
		if raw != nil {
			var event *etherdata.SuccessfullyMintedByNftEvent
			event, err = l.converter.SuccessfulMintByNft(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert raw event to the needed format")
			}

			ch <- *event
		}
	}

	return nil
}

func (l *tokenListener) readSuccessfulMintByNftEvents(ch chan<- etherdata.SuccessfullyMintedByNftEvent) (err error) {
	lastChainBlock, err := l.rpc.BlockNumber(l.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get last block in chain")
	}

	if l.maxDepth == nil {
		return l.readSuccessfulMintByNftInterval(helpers.Interval{
			From: *l.from,
			To:   lastChainBlock,
		}, ch)
	}

	if l.to == nil {
		l.to = &lastChainBlock
	}

	for _, interval := range helpers.SplitIntoIntervals(*l.from, *l.to, *l.maxDepth) {
		if err = l.readSuccessfulMintByNftInterval(interval, ch); err != nil {
			return errors.Wrap(err, "failed to read mint by nft interval", logan.F{
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

func (l *tokenListener) listenSuccessfulMintByNftEvents(ch chan<- etherdata.SuccessfullyMintedByNftEvent) (err error) {
	opts := bind.WatchOpts{
		Context: l.ctx,
	}

	filterer, err := l.getWSInstance(*l.address)
	if err != nil {
		return errors.Wrap(err, "failed to initialize a filterer")
	}

	eventsChannel := make(chan *tokencontract.TokencontractSuccessfullyMintedByNFT)
	subscription, err := filterer.WatchSuccessfullyMintedByNFT(&opts, eventsChannel, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to watch successfully minted by nft events")
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

			var convertedEvent *etherdata.SuccessfullyMintedByNftEvent
			convertedEvent, err = l.converter.SuccessfulMintByNft(*raw)
			if err != nil {
				return errors.Wrap(err, "failed to convert event to the needed type")
			}

			ch <- *convertedEvent
		}
	}
}

func (l *tokenListener) WatchSuccessfulMintByNftEvents(ch chan<- etherdata.SuccessfullyMintedByNftEvent) (err error) {
	// We firstly range from l.from to the last chain block and then start a listener

	if err = l.validateParameters(); err != nil {
		return err
	}
	if err = l.readSuccessfulMintByNftEvents(ch); err != nil {
		return errors.Wrap(err, "failed to read previous events")
	}
	if err = l.listenSuccessfulMintByNftEvents(ch); err != nil {
		return errors.Wrap(err, "failed to listen to deploy events")
	}

	return nil
}
