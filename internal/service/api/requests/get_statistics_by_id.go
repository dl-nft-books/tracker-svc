package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetStatisticsByBookRequest struct {
	pgdb.OffsetPageParams
	BookId int64 `json:"-"`
}

func NewGetStatisticsByBookRequest(r *http.Request) (*GetStatisticsByBookRequest, error) {
	var request GetStatisticsByBookRequest
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, validation.Errors{"id": errors.New("id should be a valid int64")}
	}
	request.BookId = id
	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return &request, nil
}
