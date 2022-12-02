package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type (
	// FactoryIterator is an interface that implements functions
	// needed to read factory contract events through number of blocks
	// process them simultaneously
	FactoryIterator interface {
		// From is a method specifying the block to begin with
		From(from uint64) FactoryIterator
		// To is a method specifying the block to end with
		To(to uint64) FactoryIterator
		// WithAddress is a method specifying the address factory
		// listener should be looking at
		WithAddress(address common.Address) FactoryIterator
		// WithCtx is a method modifying a FactoryReader struct to
		// use specified context
		WithCtx(ctx context.Context) FactoryIterator
		// WithProcessor is a method adding a processor to
		// a listener that will be used when catching events
		WithProcessor(processor FactoryProcessor) FactoryIterator

		// ProcessDeployEvents is a method that will
		// run between blocks specified in From and To methods
		// and process all caught contract deploy events
		ProcessDeployEvents() error
	}

	// TokenIterator is an interface that implements functions
	// needed to read factory contract events through number of blocks
	// process them simultaneously
	TokenIterator interface {
		// From is a method specifying the block to begin with
		From(from uint64) TokenIterator
		// To is a method specifying the block to end with
		To(to uint64) TokenIterator
		// WithAddress is a method specifying the address factory
		// listener should be looking at
		WithAddress(address common.Address) TokenIterator
		// WithCtx is a method modifying a FactoryReader struct to
		// use specified context
		WithCtx(ctx context.Context) TokenIterator
		// WithProcessor is a method adding a processor to
		// a listener that will be used when catching events
		WithProcessor(processor TokenProcessor) TokenIterator

		// ProcessSuccessfulMints is a method that will
		// run between blocks specified in From and To methods
		// and process all caught successful mint events
		ProcessSuccessfulMints() error
		// ProcessTransfers is a method that will
		// run between blocks specified in From and To methods
		// and process all caught transfer events
		ProcessTransfers() error
		// ProcessUpdates is a method that will
		// run between blocks specified in From and To methods
		// and process all caught update events
		ProcessUpdates() error
	}
)
