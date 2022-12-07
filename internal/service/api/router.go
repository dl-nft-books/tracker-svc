package api

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/postgres"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (a *api) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(a.log),
		ape.LoganMiddleware(a.log),
		ape.CtxMiddleware(
			// base configs
			handlers.CtxLog(a.log),
			handlers.CtxDB(postgres.NewDB(a.cfg.DB())),
			// connectors
			handlers.CtxBooker(a.cfg.BookerConnector()),
		),
	)

	// TODO: REPLACE MAIN ENDPOINT TO `runners`
	r.Route("/integrations/token-tracker", func(r chi.Router) {
		r.Route("/payments", func(r chi.Router) {
			r.Get("/", handlers.ListPayments)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentById)
			})
		})
	})

	return r
}
