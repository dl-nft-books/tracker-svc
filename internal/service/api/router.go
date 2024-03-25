package api

import (
	"github.com/dl-nft-books/tracker-svc/internal/data/postgres"
	"github.com/dl-nft-books/tracker-svc/internal/service/api/handlers"

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

	r.Route("/integrations/tracker", func(r chi.Router) {
		r.Route("/payments", func(r chi.Router) {
			r.Get("/", handlers.ListPayments)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPaymentById)
			})
		})
		r.Route("/statistics", func(r chi.Router) {
			r.Get("/", handlers.GetStatistics)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetStatisticsByBook)
			})
		})
	})

	return r
}
