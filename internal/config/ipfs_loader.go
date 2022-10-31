package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
)

const (
	ipfsLoaderYamlKey = "ipfs_loader"

	infuraMode = "infura"
	pinataMode = "pinata"
)

var ModeNotIdentified = errors.New("specified mode does not exist")

type IpfsLoader struct {
	Mode string `fig:"mode,required"`
}

func (c *config) IpfsLoader() ipfs_loader.LoaderImplementation {
	return c.ipfsLoaderOnce.Do(func() interface{} {
		var cfg IpfsLoader

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, ipfsLoaderYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		switch cfg.Mode {
		case pinataMode:
			return c.PinataImplementation()
		case infuraMode:
			return c.InfuraImplementation()
		default:
			panic(errors.From(ModeNotIdentified, logan.F{
				"mode": cfg.Mode,
			}))
		}
	}).(ipfs_loader.LoaderImplementation)
}
