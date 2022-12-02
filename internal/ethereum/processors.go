package ethereum

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
)

type (
	// FactoryProcessor is an interface that contains all functions
	// that process supported events on the factory contract
	FactoryProcessor interface {
		// ProcessDeploy processes a contract deployed event
		ProcessDeploy(event etherdata.ContractDeployedEvent) error
	}

	// TokenProcessor is an interface that contains all functions
	// that process supported events on the token contract
	TokenProcessor interface {
		// ProcessSuccessfulMint processes a successful mint event
		ProcessSuccessfulMint(event etherdata.SuccessfulMintEvent) error
		// ProcessTransfer processes a transfer event
		ProcessTransfer(event etherdata.TransferEvent) error
		// ProcessUpdate processes an update event
		ProcessUpdate(event etherdata.UpdateEvent) error
	}
)
