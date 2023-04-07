package handlers

import (
	booker "github.com/dl-nft-books/book-svc/connector"
	"github.com/dl-nft-books/book-svc/connector/models"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func ListPayments(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListPaymentsRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	qPayments, err := applyQFiltersPayments(DB(r).Payments(), Booker(r), request)
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

func applyQFiltersPayments(qPayments data.PaymentsQ, booker *booker.Connector, request *requests.ListPaymentsRequest) (*data.PaymentsQ, error) {
	if len(request.Id) > 0 {
		qPayments = qPayments.FilterById(request.Id...)
	}
	if len(request.TokenAddress) > 0 {
		qPayments = qPayments.FilterByTokenAddress(request.TokenAddress...)
	}
	if len(request.BookId) > 0 {
		booksResponse, err := booker.ListBooks(models.ListBooksParams{
			Id:      request.BookId,
			ChainId: request.ChainId,
		})

		if err != nil {
			return nil, errors.Wrap(err, "failed to select books based on book ids")
		}

		bookContracts := make([]string, 0)
		for _, book := range booksResponse.Data {
			for _, network := range book.Attributes.Networks {
				bookContracts = append(bookContracts, network.Attributes.ContractAddress)
			}
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
		paymentResource, err := payment.Resource()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert payment to the resource format")
		}

		pairDataRelationships, err := getPaymentRelationships(payment, Booker(r))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get payment relationships")
		}

		paymentResource.Relationships = pairDataRelationships

		response.Data = append(response.Data, *paymentResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
