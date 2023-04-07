package handlers

import (
	"net/http"

	booker "github.com/dl-nft-books/book-svc/connector"
	"github.com/dl-nft-books/book-svc/connector/models"
	"github.com/spf13/cast"

	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetNftPaymentById(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPaymentRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	payment, err := DB(r).NftPayments().New().FilterById(request.Id).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if payment == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	paymentResponse, err := getNftPaymentResponse(*payment, Booker(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to get payment response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *paymentResponse)
}

func getNftPaymentResponse(payment data.NftPayment, booker *booker.Connector) (*resources.NftPaymentResponse, error) {
	var paymentResponse resources.NftPaymentResponse

	pairRelationships, err := getNftPaymentRelationships(payment, booker)
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

func getNftPaymentRelationships(payment data.NftPayment, booker *booker.Connector) (*resources.NftPaymentRelationships, error) {
	bookId, err := getBookIdFromNftPayment(payment, booker)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book id from payment")
	}

	bookKey := resources.NewKeyInt64(*bookId, resources.BOOK)

	return &resources.NftPaymentRelationships{
		Book: &resources.Relation{
			Data: &bookKey,
		},
	}, nil
}

func getBookIdFromNftPayment(payment data.NftPayment, booker *booker.Connector) (*int64, error) {
	bookResponse, err := booker.ListBooks(models.ListBooksParams{
		Contract: []string{payment.ContractAddress},
		ChainId:  []int64{payment.ChainId},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to get book via connector")
	}
	if bookResponse == nil {
		return nil, errors.From(BookNotFoundErr, logan.F{
			"book_contract_address": payment.ContractAddress,
		})
	}

	bookId := cast.ToInt64(bookResponse.Data[0].Key.ID)
	return &bookId, nil
}
