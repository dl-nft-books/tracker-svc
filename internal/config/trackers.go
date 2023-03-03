package config

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const trackersYamlKey = "trackers"

type Trackers struct {
	Prefix                string          `fig:"prefix"`
	MaxDepth              uint64          `fig:"max_depth"`
	DelayBetweenIntervals time.Duration   `fig:"delay_between_intervals"`
	Backoff               BackoffSettings `fig:"backoff_settings"`
}

var defaultTrackers = Trackers{
	Prefix:                "tracker",
	MaxDepth:              5000,
	DelayBetweenIntervals: 500 * time.Millisecond,
	Backoff:               defaultBackoffSettings,
}

func (c *config) Trackers() Trackers {
	return c.trackersOnce.Do(func() interface{} {
		cfg := defaultTrackers

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks, contractHook).
			From(kv.MustGetStringMap(c.getter, trackersYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out mint runners config"))
		}

		return cfg
	}).(Trackers)
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
