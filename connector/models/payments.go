package models

import (
	"github.com/dl-nft-books/tracker-svc/internal/service/api/requests"
	"github.com/dl-nft-books/tracker-svc/resources"
)

type (
	// ListPaymentsParams is a helper struct to be included when calling ListPayments request
	ListPaymentsParams    requests.ListPaymentsRequest
	ListNftPaymentsParams requests.ListNftPaymentsRequest

	GetPaymentResponse      resources.PaymentResponse
	ListPaymentsResponse    resources.PaymentListResponse
	GetNftPaymentResponse   resources.NftPaymentResponse
	ListNftPaymentsResponse resources.NftPaymentListResponse
)
