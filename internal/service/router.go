package service

import (
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
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
