package tokend_uploader

import (
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs"
)

type TokenDIpfsUploader interface {
	TokenDIpfsImplementation() ipfs.Uploader
}

type tokend_uploader struct {
	getter kv.Getter
}

func NewTokenDIpfsUploader(getter kv.Getter) TokenDIpfsUploader {
	return &tokend_uploader{getter: getter}
}

func (t *tokend_uploader) TokenDIpfsImplementation() ipfs.Uploader {
	return NewTokenDLoader(ipfsApi.NewShell("https://ipfs.tokend.io/api"))
}
