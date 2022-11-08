package handlers

import (
	"net/http"

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

	payment, err := TrackerDB(r).Payments().New().FilterById(request.Id).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if payment == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	paymentResponse, err := getPaymentResponse(*payment, TrackerDB(r), BooksQ(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *paymentResponse)
}

func getPaymentResponse(payment data.Payment, trackerDB data.TrackerDB, qBooks data.BookQ) (*resources.PaymentResponse, error) {
	var paymentResponse resources.PaymentResponse

	pairRelationships, err := getPaymentRelationships(payment, trackerDB, qBooks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pair relationship")
	}

	paymentResource := payment.Resource()
	paymentResource.Relationships = pairRelationships

	return &paymentResponse, nil
}

func getPaymentRelationships(payment data.Payment, trackerDB data.TrackerDB, qBooks data.BookQ) (*resources.PaymentRelationships, error) {
	bookId, err := getBookIdFromPayment(payment, trackerDB, qBooks)
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

func getBookIdFromPayment(payment data.Payment, trackerDB data.TrackerDB, qBooks data.BookQ) (*int64, error) {
	contract, err := trackerDB.Contracts().New().Get(payment.ContractId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book contract from the database")
	}
	if contract == nil {
		return nil, errors.From(ContractNotFoundErr, logan.F{
			"contract_id": payment.ContractId,
		})
	}

	book, err := qBooks.FilterByContractAddress(contract.Contract).FilterActual().Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book from the database")
	}
	if book == nil {
		return nil, errors.From(BookNotFoundErr, logan.F{
			"book_contract_address": contract.Contract,
		})
	}

	return &book.ID, nil
}
