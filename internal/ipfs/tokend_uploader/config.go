package tokend_uploader

import (
	"github.com/dl-nft-books/tracker-svc/internal/ipfs"
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/kit/kv"
)

type TokenDIpfsUploader interface {
	TokenDIpfsImplementation() ipfs.Uploader
}

type tokendUploader struct {
	getter kv.Getter
}

func NewTokenDIpfsUploader(getter kv.Getter) TokenDIpfsUploader {
	return &tokendUploader{getter: getter}
}

func (t *tokendUploader) TokenDIpfsImplementation() ipfs.Uploader {
	return NewTokenDLoader(ipfsApi.NewShell("https://ipfs.tokend.io/api/"))
}
