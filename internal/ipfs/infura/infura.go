package infura

import (
	"bytes"
	"fmt"

	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
)

type InfuraLoader struct {
	api *ipfsApi.Shell
}

func NewInfuraLoader(shell *ipfsApi.Shell) ipfs.Uploader {
	return &InfuraLoader{
		api: shell,
	}
}

func (l *InfuraLoader) Upload(_ string, file []byte) (*string, error) {
	cid, err := l.api.Add(bytes.NewReader(file))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add file to the ipfs")
	}

	response := fmt.Sprintf("File's hash is %s", cid)
	return &response, nil
}
