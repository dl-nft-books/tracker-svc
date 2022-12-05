package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs/infura"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs/pinata"
	generatorConnector "gitlab.com/tokend/nft-books/generator-svc/connector/config"
)

type Config interface {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser

	// IPFS loaders
	infura.Infurer
	pinata.Pinater
	IpfsLoader() ipfs.Uploader

	// Contract trackers
	EtherClient() EtherClient
	NativeToken() etherdata.Erc20Info

	FactoryTracker() FactoryTracker
	ContractTracker() ContractTracker
	TransferTracker() ContractTracker
	MintTracker() ContractTracker
	UpdateTracker() ContractTracker

	// Connectors
	documenter.Documenter
	booker.Booker
	generatorConnector.GeneratorConfigurator
}

type config struct {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser

	// IPFS loaders
	infura.Infurer
	pinata.Pinater
	ipfsLoaderOnce comfig.Once

	// Contract trackers
	getter kv.Getter

	mintTrackerOnce     comfig.Once
	contractTrackerOnce comfig.Once
	transferTrackerOnce comfig.Once
	factoryTrackerOnce  comfig.Once
	nativeTokenOnce     comfig.Once
	ethererOnce         comfig.Once
	updateTrackerOnce   comfig.Once

	// Connectors
	booker.Booker
	generatorConnector.GeneratorConfigurator
	documenter.Documenter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Databaser:             pgdb.NewDatabaser(getter),
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter:            documenter.NewDocumenter(getter),
		Infurer:               infura.NewInfurer(getter),
		Pinater:               pinata.NewPinater(getter),
		Booker:                booker.NewBooker(getter),
		GeneratorConfigurator: generatorConnector.NewGeneratorConfigurator(getter),
	}
}
