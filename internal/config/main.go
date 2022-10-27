package config

import (
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/s3_connector"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	Databaser

	IpfsApi() *ipfsApi.Shell
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
	Databaser

	getter             kv.Getter
	mintTrackerOnce    comfig.Once
	factoryTrackerOnce comfig.Once
	ethererOnce        comfig.Once
	ipfsConnectorOnce  comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Documenter: s3_connector.NewDocumenter(getter),
	}
}
