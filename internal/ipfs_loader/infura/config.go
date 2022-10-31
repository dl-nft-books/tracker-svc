package infura

import (
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/ipfs_loader"
	"net/http"
)

const infuraConnectorYamlKey = "infura_connector"

type InfuraConfig struct {
	Endpoint      string `fig:"endpoint"`
	ProjectId     string `fig:"project_id"`
	ProjectSecret string `fig:"project_secret"`
}

type Infurer interface {
	InfuraImplementation() ipfs_loader.LoaderImplementation
}

type infurer struct {
	getter kv.Getter
}

func NewInfurer(getter kv.Getter) Infurer {
	return &infurer{getter: getter}
}

func (i *infurer) InfuraImplementation() ipfs_loader.LoaderImplementation {
	mapData := kv.MustGetStringMap(i.getter, infuraConnectorYamlKey)

	var cfg InfuraConfig
	err := figure.Out(&cfg).With(figure.BaseHooks).From(mapData).Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to read pinata connector from config"))
	}

	shell := ipfsApi.NewShellWithClient(cfg.Endpoint, newIpfsClient(cfg.ProjectId, cfg.ProjectSecret))
	return NewInfuraLoader(shell)
}

func newIpfsClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}
