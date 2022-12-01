package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	booker "gitlab.com/tokend/nft-books/book-svc/connector/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs-uploader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs-uploader/infura"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs-uploader/pinata"
	generatorConnector "gitlab.com/tokend/nft-books/generator-svc/connector/config"
)

type Config interface {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	// IPFS loaders
	infura.Infurer
	pinata.Pinater
	IpfsLoader() ipfs_uploader.Uploader

	// Contract trackers
	EtherClient() EtherClient
	NativeToken() ethereum.Erc20Info

	FactoryTracker() FactoryTracker
	TransferTracker() ContractTracker
	MintTracker() ContractTracker
	UpdateTracker() ContractTracker

	// Connectors
	documenter.Documenter
	booker.Booker
	generatorConnector.GeneratorConfigurator

	// Database
	Databaser
}

type config struct {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	// IPFS loaders
	infura.Infurer
	pinata.Pinater
	ipfsLoaderOnce comfig.Once

	// Contract trackers
	mintTrackerOnce     comfig.Once
	transferTrackerOnce comfig.Once
	factoryTrackerOnce  comfig.Once
	nativeTokenOnce     comfig.Once
	ethererOnce         comfig.Once
	updateTrackerOnce   comfig.Once

	// Connectors
	booker.Booker
	generatorConnector.GeneratorConfigurator
	documenter.Documenter

	// Database
	Databaser

	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Databaser:             NewDatabaser(getter),
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
