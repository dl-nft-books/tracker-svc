package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

func ListPayments(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListPaymentsRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	qPayments, err := applyQFiltersPayments(TrackerDB(r).Payments(), BooksQ(r), request)
	if err != nil {
		Log(r).WithError(err).Error("unable to apply request filters")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	q := *qPayments
	payments, err := q.Select()
	if err != nil {
		Log(r).WithError(err).Error("unable to select payments from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response, err := formListPaymentsResponse(r, request, payments)
	if err != nil {
		Log(r).WithError(err).Error("unable to form payments list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, response)
}

func applyQFiltersPayments(qPayments data.PaymentsQ, qBooks data.BookQ, request *requests.ListPaymentsRequest) (*data.PaymentsQ, error) {
	if len(request.Id) > 0 {
		qPayments = qPayments.FilterById(request.Id...)
	}
	if len(request.TokenAddress) > 0 {
		qPayments = qPayments.FilterByTokenAddress(request.TokenAddress...)
	}
	if len(request.BookId) > 0 {
		books, err := qBooks.FilterByID(request.BookId...).FilterActual().Select()
		if err != nil {
			return nil, errors.Wrap(err, "failed to select books based on book ids")
		}

		bookContracts := make([]string, 0)
		for _, book := range books {
			bookContracts = append(bookContracts, book.ContractAddress)
		}

		qPayments = qPayments.FilterByContractAddress(bookContracts...)
	}

	qPayments = qPayments.Sort(request.Sorts)
	qPayments = qPayments.Page(request.OffsetPageParams)

	return &qPayments, nil
}

func formListPaymentsResponse(r *http.Request, request *requests.ListPaymentsRequest, paymentsList []data.Payment) (*resources.PaymentListResponse, error) {
	var response resources.PaymentListResponse
	if len(paymentsList) == 0 {
		return &resources.PaymentListResponse{
			Data: make([]resources.Payment, 0),
		}, nil
	}

	for _, payment := range paymentsList {
		paymentResource := payment.Resource()
		pairDataRelationships, err := getPaymentRelationships(payment, TrackerDB(r), BooksQ(r))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get payment relationships")
		}

		paymentResource.Relationships = pairDataRelationships

		response.Data = append(response.Data, paymentResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
