package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	consumersYamlKey = "consumers"
	routinerYamlKey  = "routiner"
)

type Runner struct {
	Prefix  string          `fig:"prefix"`
	Backoff BackoffSettings `fig:"backoff"`
}

func (c *config) runner(once *comfig.Once, yamlKey string) Runner {
	return once.Do(func() interface{} {
		cfg := Runner{
			Prefix:  yamlKey,
			Backoff: defaultBackoffSettings,
		}

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, yamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out contract runners config"))
		}

		return cfg
	}).(Runner)
}

func (c *config) Consumers() Runner {
	return c.runner(&c.consumersOnce, consumersYamlKey)
}

func (c *config) Routiner() Runner {
	return c.runner(&c.routinerOnce, routinerYamlKey)
}
