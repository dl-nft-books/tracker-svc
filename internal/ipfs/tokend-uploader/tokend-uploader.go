package tokend_uploader

import (
	"bytes"
	"fmt"
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
)

type IpfsUploader struct {
	api *ipfsApi.Shell
}

func NewTokenDLoader(shell *ipfsApi.Shell) ipfs.Uploader {
	return &IpfsUploader{
		api: shell,
	}
}

func (l *IpfsUploader) Upload(_ string, file []byte) (*string, error) {
	fileHash, err := l.api.Add(bytes.NewReader(file))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add file to the ipfs")
	}
	response := fmt.Sprintf("File's hash is %s", fileHash)
	return &response, nil
}
