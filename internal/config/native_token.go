package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/etherdata"
)

const nativeTokenYamlKey = "native_token"

func (c *config) NativeToken() etherdata.Erc20Info {
	return c.nativeTokenOnce.Do(func() interface{} {
		var nativeTokenInfo etherdata.Erc20Info

		if err := figure.
			Out(&nativeTokenInfo).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, nativeTokenYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out native token info"))
		}

		// This is probably set by default, but just to make sure
		nativeTokenInfo.TokenAddress = etherdata.NullAddress

		return nativeTokenInfo
	}).(etherdata.Erc20Info)
}
