package pinata

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/tracker-svc/internal/ipfs"
)

const pinataConnectorYamlKey = "pinata_connector"

type Pinater interface {
	PinataImplementation() ipfs.Uploader
}

type PinataConfig struct {
	Endpoint  string `fig:"endpoint"`
	ApiKey    string `fig:"api_key"`
	ApiSecret string `fig:"api_secret"`
}

type pinater struct {
	getter kv.Getter
}

func NewPinater(getter kv.Getter) Pinater {
	return &pinater{getter: getter}
}

func (p *pinater) PinataImplementation() ipfs.Uploader {
	mapData := kv.MustGetStringMap(p.getter, pinataConnectorYamlKey)

	var cfg PinataConfig
	err := figure.Out(&cfg).With(figure.BaseHooks).From(mapData).Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to read pinata connector from config"))
	}

	return NewPinataLoader(&cfg)
}
