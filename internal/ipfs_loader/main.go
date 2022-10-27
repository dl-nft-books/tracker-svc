package ipfs_loader

import (
	"bytes"
	"fmt"
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/s3_connector"
	"io/ioutil"
	"net/http"
	"strings"
)

type IpfsLoader struct {
	documenter *s3_connector.Connector
	api        *ipfsApi.Shell
	logger     *logan.Entry
}

func NewIpfsLoader(cfg config.Config) *IpfsLoader {
	return &IpfsLoader{
		documenter: cfg.DocumenterConnector(),
		api:        cfg.IpfsApi(),
		logger:     cfg.Log(),
	}
}

func (l *IpfsLoader) Load(uri string) error {
	l.logger.Debugf("Caught request to process uri %s", uri)

	fileName := fmt.Sprintf("%s.pdf", l.GetHashOutUri(uri))

	l.logger.Debugf("Trying to find file %s", fileName)

	documentLink, err := l.documenter.GetDocumentLink(fileName)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	documentUrl := documentLink.Data.Attributes.Url
	l.logger.Debugf("Trying to download via %s", documentUrl)

	downloadedDocument, err := downloadDocument(documentUrl)
	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}

	cid, err := l.api.Add(bytes.NewReader(downloadedDocument))
	if err != nil {
		return errors.Wrap(err, "failed to add file to the ipfs")
	}
	l.logger.Debugf("Successfully loaded document with cid %s", cid)

	_, err = l.documenter.DeleteDocument(fileName)
	if err != nil {
		return errors.Wrap(err, "failed to delete document")
	}
	l.logger.Debug("Successfully removed document")

	return nil
}

func (l *IpfsLoader) GetHashOutUri(uri string) (hash string) {
	split := strings.Split(uri, "/")
	return split[len(split)-1]
}

func downloadDocument(link string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
