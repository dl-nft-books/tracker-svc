package config

import (
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
	"reflect"
)

const nativeTokenYamlKey = "native_token"

func (c *config) NativeToken() etherdata.Erc20Info {
	return c.nativeTokenOnce.Do(func() interface{} {
		var nativeTokenInfo etherdata.Erc20Info

		if err := figure.
			Out(&nativeTokenInfo).
			With(figure.BaseHooks, uint8Hook).
			From(kv.MustGetStringMap(c.getter, nativeTokenYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out native token info"))
		}

		// This is probably set by default, but just to make sure
		nativeTokenInfo.TokenAddress = etherdata.NullAddress

		return nativeTokenInfo
	}).(etherdata.Erc20Info)
}

var uint8Hook = figure.Hooks{
	"uint8": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToUint8E(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse int")
		}
		return reflect.ValueOf(result), nil
	},
}
