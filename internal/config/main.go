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
	generatorer "gitlab.com/tokend/nft-books/generator-svc/connector"
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

	Trackers() Trackers
	Consumers() Runner
	Routiner() Runner

	// Connectors
	documenter.Documenter
	booker.Booker
	generatorer.GeneratorConfigurator
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
	getter          kv.Getter
	trackersOnce    comfig.Once
	consumersOnce   comfig.Once
	routinerOnce    comfig.Once
	nativeTokenOnce comfig.Once
	ethererOnce     comfig.Once

	// Connectors
	booker.Booker
	generatorer.GeneratorConfigurator
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
		GeneratorConfigurator: generatorer.NewGeneratorConfigurator(getter),
	}
}
