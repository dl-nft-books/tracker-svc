package config

import (
	documenter "github.com/dl-nft-books/blob-svc/connector/config"
	booker "github.com/dl-nft-books/book-svc/connector"
	core "github.com/dl-nft-books/core-svc/connector"
	networker "github.com/dl-nft-books/network-svc/connector"
	"github.com/dl-nft-books/tracker-svc/internal/ipfs/infura"
	"github.com/dl-nft-books/tracker-svc/internal/ipfs/pinata"
	tokendUploader "github.com/dl-nft-books/tracker-svc/internal/ipfs/tokend_uploader"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
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
	tokendUploader.TokenDIpfsUploader
	IpfsLoader() IpfsLoader

	Trackers() Trackers
	Consumers() Runner
	Routiner() Runner

	// Connectors
	documenter.Documenter
	booker.Booker
	networker.NetworkConfigurator
	core.GeneratorConfigurator
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
	tokendUploader.TokenDIpfsUploader
	ipfsLoaderOnce comfig.Once

	// Addr trackers
	getter        kv.Getter
	trackersOnce  comfig.Once
	consumersOnce comfig.Once
	routinerOnce  comfig.Once

	// Connectors
	booker.Booker
	networker.NetworkConfigurator
	core.GeneratorConfigurator
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
		TokenDIpfsUploader:    tokendUploader.NewTokenDIpfsUploader(getter),
		Booker:                booker.NewBooker(getter),
		NetworkConfigurator:   networker.NewNetworkConfigurator(getter),
		GeneratorConfigurator: core.NewGeneratorConfigurator(getter),
	}
}
