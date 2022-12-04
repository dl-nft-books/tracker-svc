package api

import (
	"context"
	"net"
	"net/http"

	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners"

	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	cfg      config.Config
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	ctx := context.Background()

	if err := runners.Run(s.cfg, ctx); err != nil {
		return errors.Wrap(err, "failed to initialize runners")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		cfg:      cfg,
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(errors.Wrap(err, "failed to run new service"))
	}
}
