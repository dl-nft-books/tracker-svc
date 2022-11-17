package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/ethereum"
)

const nativeTokenYamlKey = "native_token"

func (c *config) NativeToken() ethereum.Erc20Info {
	return c.nativeTokenOnce.Do(func() interface{} {
		var nativeTokenInfo ethereum.Erc20Info

		if err := figure.
			Out(&nativeTokenInfo).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, nativeTokenYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out native token info"))
		}

		// This is probably set by default, but just to make sure
		nativeTokenInfo.TokenAddress = ethereum.NullAddress

		return nativeTokenInfo
	}).(ethereum.Erc20Info)
}
