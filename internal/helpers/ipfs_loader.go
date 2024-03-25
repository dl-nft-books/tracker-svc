package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dl-nft-books/tracker-svc/internal/data/opensea"

	documenter "github.com/dl-nft-books/blob-svc/connector/api"
	"github.com/dl-nft-books/tracker-svc/internal/config"
	"github.com/dl-nft-books/tracker-svc/internal/ipfs"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type IpfsLoader struct {
	implementation ipfs.Uploader
	documenter     *documenter.Connector
	logger         *logan.Entry
	BaseUri        string
}

func NewIpfsLoader(cfg config.Config) *IpfsLoader {
	return &IpfsLoader{
		implementation: cfg.IpfsLoader().Uploader,
		documenter:     cfg.DocumenterConnector(),
		logger:         cfg.Log(),
		BaseUri:        cfg.IpfsLoader().BaseUri,
	}
}

func (l *IpfsLoader) UploadBanner(uri string) error {
	bannerName := fmt.Sprintf("%s.png", uri)

	l.logger.Debugf("Caught request to process uri %s", uri)
	l.logger.Debugf("Trying to find banner %s", bannerName)

	documentLink, err := l.documenter.GetDocumentLink(bannerName)
	if err != nil {
		return errors.Wrap(err, "failed to get document link")
	}

	documentUrl := documentLink.Data.Attributes.Url
	l.logger.Debugf("Trying to download via %s", documentUrl)

	downloadedDocument, err := downloadDocument(documentUrl)

	if err != nil {
		return errors.Wrap(err, "failed to download document")
	}

	l.logger.Debug("Downloaded document. Trying to load it on the IPFS")

	response, err := l.implementation.Upload(uri, downloadedDocument)
	if err != nil {
		return errors.Wrap(err, "failed to add banner to the ipfs")
	}

	l.logger.Debugf("Successfully loaded document.\nResponse: %s", *response)

	if _, err = l.documenter.DeleteDocument(bannerName); err != nil {
		return errors.Wrap(err, "failed to delete document")
	}

	l.logger.Debug("Successfully removed document from S3")
	l.logger.Debug("Request was successfully processed")

	return nil
}

func (l *IpfsLoader) UploadMetadata(info opensea.Metadata) error {
	l.logger.Debugf("Caught request to process metadata")
	l.logger.Debug("Marshalling metadata...")

	raw, err := json.Marshal(info)
	if err != nil {
		return errors.Wrap(err, "failed to marshal metadata")
	}

	l.logger.Debug("Metadata marshalled successfully")
	l.logger.Debug("Loading metadata to IPFS...")

	metadataBannerName := l.GetHashOutUri(info.Image) + "-meta"

	response, err := l.implementation.Upload(metadataBannerName, raw)
	if err != nil {
		return errors.Wrap(err, "failed to add metadata to the ipfs")
	}

	l.logger.Debugf("Metadata was successfully loaded to IPFS.\nResponse: %s", *response)
	l.logger.Debug("Request was successfully processed")

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
