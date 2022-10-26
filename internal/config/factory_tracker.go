package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
	"time"
)

const factoryTrackerYamlKey = "factory_tracker"

type Runner struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"normal_period"`
}

type FactoryTracker struct {
	Name          string         `fig:"name"`
	Address       common.Address `fig:"address"`
	Runner        Runner         `fig:"runner"`
	FirstBlock    uint64         `fig:"first_block"`
	IterationSize uint64         `fig:"iteration_size"`
}

var defaultFactoryTracker = FactoryTracker{
	Name:    "factory_tracker",
	Address: common.Address{},
	Runner: Runner{
		NormalPeriod:      time.Minute,
		MinAbnormalPeriod: time.Minute,
		MaxAbnormalPeriod: time.Minute,
	},
	FirstBlock:    0,
	IterationSize: 100,
}

func (c *config) FactoryTracker() FactoryTracker {
	return c.factoryTrackerOnce.Do(func() interface{} {
		cfg := defaultFactoryTracker

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, contractHook).
			From(kv.MustGetStringMap(c.getter, factoryTrackerYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint tracker config"))
		}

		return cfg
	}).(FactoryTracker)
}

var contractHook = figure.Hooks{
	"common.Address": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			address := common.HexToAddress(v)
			return reflect.ValueOf(address), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
