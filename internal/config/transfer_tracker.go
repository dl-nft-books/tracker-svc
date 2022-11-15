package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

const transferTrackerYamlKey = "transfer_tracker"

type TransferTracker struct {
	Name          string `fig:"name"`
	Capacity      int64  `fig:"capacity"`
	IterationSize uint64 `fig:"iteration_size"`
	Runner        Runner `fig:"runner"`
}

var defaultTransferTracker = TransferTracker{
	Name:          "transfer_tracker",
	Capacity:      1,
	IterationSize: 100,
	Runner:        defaultRunner,
}

func (c *config) TransferTracker() TransferTracker {
	return c.transferTrackerOnce.Do(func() interface{} {
		cfg := defaultTransferTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, transferTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out transfer tracker config"))
		}

		return cfg
	}).(TransferTracker)
}
