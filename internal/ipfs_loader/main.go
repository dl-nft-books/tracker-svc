package ipfs_loader

import "gitlab.com/tokend/nft-books/contract-tracker/internal/config"

type IpfsLoader struct {
}

func NewIpfsLoader(cfg config.Config) {

}

func (l *IpfsLoader) Load(hash string) error {
	return nil
}
