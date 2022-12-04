package api

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			// base configs
			handlers.CtxLog(s.log),
			handlers.CtxDB(postgres.NewDB(s.cfg.DB())),
			// connectors
			handlers.CtxBooker(s.cfg.BookerConnector()),
			handlers.CtxGeneratorer(*s.cfg.GeneratorConnector()),
		),
	)
	r.Route("/integrations/runners", func(r chi.Router) {
		r.Route("/payments", func(r chi.Router) {
			r.Get("/", handlers.ListPayments)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentById)
			})
		})
	})

	return r
}
