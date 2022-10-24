package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"reflect"
	"time"
)

const mintTrackerYamlKey = "mint_tracker"

type Runner struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"normal_period"`
}

type MintTracker struct {
	InitialBlock  uint64           `fig:"initial_block"`
	IterationSize uint64           `fig:"iteration_size"`
	Runner        Runner           `fig:"runner"`
	Contracts     []common.Address `fig:"contracts"`
}

func (c *config) MintTracker() MintTracker {
	return c.mintTrackerOnce.Do(func() interface{} {
		var cfg MintTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, mintTrackerHooks).
			From(kv.MustGetStringMap(c.getter, mintTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(MintTracker)
}

var mintTrackerHooks = figure.Hooks{
	"[]common.Address": func(value interface{}) (reflect.Value, error) {
		addressesAsStrings, err := cast.ToStringSliceE(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse []string")
		}

		result := make([]common.Address, 0)
		for _, str := range addressesAsStrings {
			result = append(result, common.HexToAddress(str))
		}

		return reflect.ValueOf(result), nil
	},
}
