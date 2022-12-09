package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
)

const ipfsLoaderYamlKey = "uploader"

var ModeNotIdentified = errors.New("specified mode does not exist")

func (c *config) implementationsMap() map[string]ipfs.Uploader {
	return map[string]ipfs.Uploader{
		"infura": c.InfuraImplementation(),
		"pinata": c.PinataImplementation(),
	}
}

type IpfsLoader struct {
	Mode string `fig:"mode,required"`
}

func (c *config) IpfsLoader() ipfs.Uploader {
	return c.ipfsLoaderOnce.Do(func() interface{} {
		var cfg IpfsLoader

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, ipfsLoaderYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint runners config"))
		}

		implementation, ok := c.implementationsMap()[cfg.Mode]
		if !ok {
			panic(errors.From(ModeNotIdentified, logan.F{
				"mode": cfg.Mode,
			}))
		}

		return implementation
	}).(ipfs.Uploader)
}
