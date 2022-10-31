package pinata

import (
	"bytes"
	"fmt"
	provider "github.com/wealdtech/go-ipfs-provider"
	pinata "github.com/wealdtech/go-ipfs-provider-pinata"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
)

var (
	NoResponseErr    = errors.New("No response received")
	FailedLoadingErr = errors.New("Failed to load a file")
	HashNotFound     = errors.New("IPFS Hash was not found in the response")
)

const (
	methodField             = "POST"
	contentTypeField        = "Content-Type"
	pinataApiKeyField       = "pinata_api_key"
	pinataApiSecretKeyField = "pinata_secret_api_key"
)

type PinataLoader struct {
	provider *pinata.Provider
}

func NewPinataLoader(cfg *PinataConfig) ipfs_loader.LoaderImplementation {
	provider, err := pinata.NewProvider(cfg.ApiKey, cfg.ApiSecret)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize a new provider"))
	}

	return &PinataLoader{
		provider: provider,
	}
}

func (l *PinataLoader) Load(name string, file []byte) (*string, error) {
	hash, err := l.provider.PinContent(name, bytes.NewReader(file), &provider.ContentOpts{StoreInDirectory: false})
	if err != nil {
		return nil, errors.Wrap(err, "failed to pin content")
	}

	response := fmt.Sprintf("File's hash is %s", hash)
	return &response, nil
}
