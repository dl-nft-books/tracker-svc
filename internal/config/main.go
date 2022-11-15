package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	s3api "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	s3config "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader/infura"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader/pinata"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	infura.Infurer
	pinata.Pinater
	Databaser

	IpfsLoader() ipfs_loader.LoaderImplementation
	DocumenterConnector() *s3api.Connector
	FactoryTracker() FactoryTracker
	TransferTracker() TransferTracker
	MintTracker() MintTracker
	EtherClient() EtherClient
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	s3config.Documenter
	infura.Infurer
	pinata.Pinater
	Databaser

	getter              kv.Getter
	mintTrackerOnce     comfig.Once
	factoryTrackerOnce  comfig.Once
	transferTrackerOnce comfig.Once
	ethererOnce         comfig.Once
	ipfsLoaderOnce      comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter: s3config.NewDocumenter(getter),
		Infurer:    infura.NewInfurer(getter),
		Pinater:    pinata.NewPinater(getter),
	}
}
