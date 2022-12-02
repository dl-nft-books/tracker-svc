package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type (
	// FactoryListener is an interface that implements functions
	// needed to listen factory contract events and process them
	// simultaneously
	FactoryListener interface {
		// WithAddress is a method specifying the address factory
		// listener should be looking at
		WithAddress(address common.Address) FactoryListener
		// WithCtx is a method modifying a FactoryListener struct to
		// use specified context
		WithCtx(ctx context.Context) FactoryListener
		// WithProcessor is a method adding a processor to
		// a listener that will be used when catching events
		WithProcessor(processor FactoryProcessor) FactoryListener

		// WatchContractCreatedEvents is a function that will run
		// a listener in a for{} loop and process all caught events
		// using a processor from WithProcessor method
		WatchContractCreatedEvents() error
	}

	// TokenListener is an interface that implements functions
	// needed to listen token contract events and process them
	// simultaneously
	TokenListener interface {
		// WithAddress is a method specifying the address factory
		// listener should be looking at
		WithAddress(address common.Address) TokenListener
		// WithCtx is a method modifying a TokenListener struct to
		// use specified context
		WithCtx(ctx context.Context) TokenListener
		// WithProcessor is a method adding a processor to
		// a listener that will be used when catching events
		WithProcessor(processor TokenProcessor) TokenListener

		// WatchSuccessfulMintEvents is a function that will run
		// a listener in a for{} loop and process all caught events
		// using a processor from WithProcessor method
		WatchSuccessfulMintEvents() error
		// WatchTransferEvents is a function that will run
		// a listener in a for{} loop and process all caught events
		// using a processor from WithProcessor method
		WatchTransferEvents() error
		// WatchUpdateEvents is a function that will run
		// a listener in a for{} loop and process all caught events
		// using a processor from WithProcessor method
		WatchUpdateEvents() error
	}
)
