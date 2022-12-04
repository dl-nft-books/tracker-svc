package ethereum

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
)

type (
	FactoryProcessor interface {
		ProcessDeploy(event etherdata.ContractDeployedEvent) error
	}

	TokenProcessor interface {
		ProcessSuccessfulMint(event etherdata.SuccessfulMintEvent) error
		ProcessTransfer(event etherdata.TransferEvent) error
		ProcessUpdate(event etherdata.UpdateEvent) error
	}
)
