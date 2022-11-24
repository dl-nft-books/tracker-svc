package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	s3api "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	s3config "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/uploader"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/uploader/infura"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/uploader/pinata"
	networkerCfg "gitlab.com/tokend/nft-books/network-svc/connector/config"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	infura.Infurer
	pinata.Pinater
	Databaser
	networkerCfg.NetworkConfigurator

	IpfsLoader() uploader.Uploader
	DocumenterConnector() *s3api.Connector
	FactoryTracker() FactoryTracker
	TransferTracker() ContractTracker
	MintTracker() ContractTracker
	NativeToken() ethereum.Erc20Info
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	s3config.Documenter
	infura.Infurer
	pinata.Pinater
	Databaser
	networkerCfg.NetworkConfigurator

	getter              kv.Getter
	mintTrackerOnce     comfig.Once
	transferTrackerOnce comfig.Once
	factoryTrackerOnce  comfig.Once
	nativeTokenOnce     comfig.Once
	ipfsLoaderOnce      comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:              getter,
		Databaser:           NewDatabaser(getter),
		Copuser:             copus.NewCopuser(getter),
		Listenerer:          comfig.NewListenerer(getter),
		Logger:              comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter:          s3config.NewDocumenter(getter),
		Infurer:             infura.NewInfurer(getter),
		Pinater:             pinata.NewPinater(getter),
		NetworkConfigurator: networkerCfg.NewNetworkConfigurator(getter),
	}
}
