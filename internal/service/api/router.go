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
			// Base configs
			handlers.CtxLog(a.log),
			handlers.CtxDB(postgres.NewDB(a.cfg.DB())),

			// Connectors
			handlers.CtxBooker(a.cfg.BookerConnector()),
		),
	)

	// TODO: Replace later 'token-tracker' to just 'tracker'
	// Not done yet since it is not synchronized with the frontend side
	r.Route("/integrations/token-tracker", func(r chi.Router) {
		r.Route("/payments", func(r chi.Router) {
			r.Get("/", handlers.ListPayments)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentById)
			})
			r.Route("/nft", func(r chi.Router) {
				r.Get("/", handlers.ListNftPayments)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handlers.GetNftPaymentById)
				})
			})
		})
	})

	return r
}
