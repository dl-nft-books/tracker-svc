package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader/infura"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader/pinata"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/s3_connector"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	infura.Infurer
	pinata.Pinater
	Databaser

	IpfsLoader() ipfs_loader.LoaderImplementation
	DocumenterConnector() *s3_connector.Connector
	FactoryTracker() FactoryTracker
	MintTracker() MintTracker
	EtherClient() EtherClient
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	s3_connector.Documenter
	infura.Infurer
	pinata.Pinater
	Databaser

	getter             kv.Getter
	mintTrackerOnce    comfig.Once
	factoryTrackerOnce comfig.Once
	ethererOnce        comfig.Once
	ipfsLoaderOnce     comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter: s3_connector.NewDocumenter(getter),
		Infurer:    infura.NewInfurer(getter),
		Pinater:    pinata.NewPinater(getter),
	}
}
