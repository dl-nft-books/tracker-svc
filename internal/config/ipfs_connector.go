package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"

	ipfsApi "github.com/ipfs/go-ipfs-api"
)

const ipfsConnectorYamlKey = "ipfs_connector"

type IpfsConfig struct {
	Endpoint      string `fig:"endpoint"`
	ProjectId     string `fig:"project_id"`
	ProjectSecret string `fig:"project_secret"`
}

func (c *config) IpfsApi() *ipfsApi.Shell {
	return c.ipfsConnectorOnce.Do(func() interface{} {
		var cfg IpfsConfig

		if err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, ipfsConnectorYamlKey)).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out ipfs connector config"))
		}

		shell := ipfsApi.NewShellWithClient(cfg.Endpoint, newIpfsClient(cfg.ProjectId, cfg.ProjectSecret))
		return shell
	}).(*ipfsApi.Shell)
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
