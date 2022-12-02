package pinata

import (
	"bytes"
	"fmt"

	provider "github.com/wealdtech/go-ipfs-provider"
	pinata "github.com/wealdtech/go-ipfs-provider-pinata"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
)

type PinataLoader struct {
	provider *pinata.Provider
}

func NewPinataLoader(cfg *PinataConfig) ipfs.Uploader {
	provider, err := pinata.NewProvider(cfg.ApiKey, cfg.ApiSecret)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize a new provider"))
	}

	return &PinataLoader{
		provider: provider,
	}
}

func (l *PinataLoader) Upload(name string, file []byte) (*string, error) {
	hash, err := l.provider.PinContent(name, bytes.NewReader(file), &provider.ContentOpts{StoreInDirectory: false})
	if err != nil {
		return nil, errors.Wrap(err, "failed to pin content")
	}

	response := fmt.Sprintf("File's hash is %s", hash)
	return &response, nil
}
