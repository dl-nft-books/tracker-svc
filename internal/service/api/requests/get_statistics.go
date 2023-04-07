package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetStatisticsRequest struct {
	NFT  pgdb.OffsetPageParams `url:"nft"`
	Sort string                `url:"nft_sort" default:"floor_price"`
}

func NewGetStatisticsRequest(r *http.Request) (*GetStatisticsRequest, error) {
	var request GetStatisticsRequest

	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return &request, nil
}
