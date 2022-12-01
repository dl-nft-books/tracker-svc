package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	mintTrackerYamlKey     = "mint_tracker"
	transferTrackerYamlKey = "transfer_tracker"
	updateTrackerYamlKey   = "update_tracker"
)

type ContractTracker struct {
	Name          string `fig:"name"`
	Capacity      uint64 `fig:"capacity"`
	IterationSize uint64 `fig:"iteration_size"`
	Runner        Runner `fig:"runner"`
}

func (c *config) contractTracker(once *comfig.Once, yamlKey string) ContractTracker {
	return once.Do(func() interface{} {
		cfg := ContractTracker{
			Name:          yamlKey,
			Capacity:      1,
			IterationSize: 100,
			Runner:        defaultRunner,
		}

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, yamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out contract tracker config"))
		}

		return cfg
	}).(ContractTracker)
}

func (c *config) MintTracker() ContractTracker {
	return c.contractTracker(&c.mintTrackerOnce, mintTrackerYamlKey)
}

func (c *config) TransferTracker() ContractTracker {
	return c.contractTracker(&c.transferTrackerOnce, transferTrackerYamlKey)
}

func (c *config) UpdateTracker() ContractTracker {
	return c.contractTracker(&c.updateTrackerOnce, updateTrackerYamlKey)
}
