package api

import (
	"net"
	"net/http"

	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type api struct {
	cfg      config.Config
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (a *api) run() error {
	a.log.Info("Service started")
	r := a.router()

	if err := a.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(a.listener, r)
}

func newService(cfg config.Config) *api {
	return &api{
		cfg:      cfg,
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) error {
	if err := newService(cfg).run(); err != nil {
		return errors.Wrap(err, "failed to run a new service")
	}

	return nil
}
