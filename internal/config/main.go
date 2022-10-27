package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	blobconnector "gitlab.com/tokend/nft-books/blob-svc/connector"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	DocumenterConnector() *blobconnector.Connector
	FactoryTracker() FactoryTracker
	MintTracker() MintTracker
	EtherClient() EtherClient
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter

	blobconnector.Documenter
	mintTrackerOnce    comfig.Once
	factoryTrackerOnce comfig.Once
	ethererOnce        comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter: blobconnector.NewDocumenter(getter),
	}
}
