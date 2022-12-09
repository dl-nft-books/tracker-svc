package handlers

import (
	"net/http"

	"github.com/spf13/cast"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/book-svc/connector/models"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var (
	ContractNotFoundErr = errors.New("contract with specified id was not found")
	BookNotFoundErr     = errors.New("book with specified address was not found")
)

func GetPaymentById(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPaymentRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	payment, err := DB(r).Payments().New().FilterById(request.Id).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if payment == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	paymentResponse, err := getPaymentResponse(*payment, DB(r), Booker(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *paymentResponse)
}

func getPaymentResponse(payment data.Payment, trackerDB data.DB, booker *booker.Connector) (*resources.PaymentResponse, error) {
	var paymentResponse resources.PaymentResponse

	pairRelationships, err := getPaymentRelationships(payment, trackerDB, booker)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pair relationship")
	}

	paymentResource, err := payment.Resource()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get payment as a resource")
	}
	paymentResource.Relationships = pairRelationships
	paymentResponse.Data = *paymentResource

	return &paymentResponse, nil
}

func getPaymentRelationships(payment data.Payment, db data.DB, booker *booker.Connector) (*resources.PaymentRelationships, error) {
	bookId, err := getBookIdFromPayment(payment, db, booker)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book id from payment")
	}

	bookKey := resources.NewKeyInt64(*bookId, resources.BOOK)

	return &resources.PaymentRelationships{
		Book: &resources.Relation{
			Data: &bookKey,
		},
	}, nil
}

func getBookIdFromPayment(payment data.Payment, db data.DB, booker *booker.Connector) (*int64, error) {
	contract, err := db.Contracts().New().Get(payment.ContractId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book contract from the database")
	}
	if contract == nil {
		return nil, errors.From(ContractNotFoundErr, logan.F{
			"contract_id": payment.ContractId,
		})
	}

	bookResponse, err := booker.ListBooks(models.ListBooksParams{
		Contract: []string{contract.Addr},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to get book via connector")
	}
	if bookResponse == nil {
		return nil, errors.From(BookNotFoundErr, logan.F{
			"book_contract_address": contract.Addr,
		})
	}

	bookId := cast.ToInt64(bookResponse.Data[0].Key.ID)
	return &bookId, nil
}
