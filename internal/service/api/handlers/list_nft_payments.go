package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/book-svc/connector/models"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
	"net/http"
)

func ListNftPayments(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListNftPaymentsRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	qPayments, err := applyQFiltersNftPayments(DB(r).NftPayments(), Booker(r), request)
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

	response, err := formListNftPaymentsResponse(r, request, payments)
	if err != nil {
		Log(r).WithError(err).Error("unable to form payments list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, response)
}

func applyQFiltersNftPayments(qPayments data.NftPaymentsQ, booker *booker.Connector, request *requests.ListNftPaymentsRequest) (*data.NftPaymentsQ, error) {
	if len(request.Id) > 0 {
		qPayments = qPayments.FilterById(request.Id...)
	}
	if len(request.NftAddress) > 0 {
		qPayments = qPayments.FilterByNftAddress(request.NftAddress...)
	}
	if len(request.NftId) > 0 {
		qPayments = qPayments.FilterByNftId(request.NftId...)
	}
	if len(request.BookId) > 0 {
		booksResponse, err := booker.ListBooks(models.ListBooksParams{
			Id: request.BookId,
		})

		if err != nil {
			return nil, errors.Wrap(err, "failed to select books based on book ids")
		}

		bookContracts := make([]string, 0)
		for _, book := range booksResponse.Data {
			bookContracts = append(bookContracts, book.Attributes.ContractAddress)
		}

		qPayments = qPayments.FilterByContractAddress(bookContracts...)
	}

	qPayments = qPayments.Sort(request.Sorts)
	qPayments = qPayments.Page(request.OffsetPageParams)

	return &qPayments, nil
}

func formListNftPaymentsResponse(r *http.Request, request *requests.ListNftPaymentsRequest, paymentsList []data.NftPayment) (*resources.NftPaymentListResponse, error) {
	var response resources.NftPaymentListResponse
	if len(paymentsList) == 0 {
		return &resources.NftPaymentListResponse{
			Data: make([]resources.NftPayment, 0),
		}, nil
	}

	for _, payment := range paymentsList {
		paymentResource, err := payment.Resource()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert payment to the resource format")
		}

		pairDataRelationships, err := getNftPaymentRelationships(payment, DB(r), Booker(r))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get payment relationships")
		}

		paymentResource.Relationships = pairDataRelationships

		response.Data = append(response.Data, *paymentResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
