package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

const mintTrackerYamlKey = "mint_tracker"

type MintTracker struct {
	Capacity      int64  `fig:"capacity"`
	IterationSize uint64 `fig:"iteration_size"`
}

var defaultMintTracker = MintTracker{
	Capacity:      1,
	IterationSize: 100,
}

func (c *config) MintTracker() MintTracker {
	return c.mintTrackerOnce.Do(func() interface{} {
		cfg := defaultMintTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, mintTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(MintTracker)
}
