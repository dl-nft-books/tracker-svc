package data

import (
	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type NftRequest struct {
	Id            int64                      `db:"id" structs:"-"`
	Requester     string                     `db:"requester" structs:"requester"`
	NftContract   string                     `db:"nft_contract" structs:"nft_contract"`
	NftId         int64                      `db:"nft_id" structs:"nft_id"`
	ChainId       int64                      `db:"chain_id" structs:"chain_id"`
	BookId        int64                      `db:"book_id" structs:"book_id"`
	Status        resources.NftRequestStatus `db:"status" structs:"status"`
	CreatedAt     time.Time                  `db:"created_at" structs:"created_at"`
	LastUpdatedAt time.Time                  `db:"last_updated_at" structs:"last_updated_at"`
}

type NftRequestsQ interface {
	New() NftRequestsQ

	Get() (*NftRequest, error)
	Select() ([]NftRequest, error)
	DeleteByID(id int64) error

	Sort(sort pgdb.Sorts) NftRequestsQ
	Page(page pgdb.OffsetPageParams) NftRequestsQ

	Insert(nftRequest NftRequest) (int64, error)
	Transaction(fn func(q NftRequestsQ) error) error
	FilterByStatus(status ...resources.NftRequestStatus) NftRequestsQ
	FilterById(id ...int64) NftRequestsQ
	FilterByChainId(id ...int64) NftRequestsQ
	FilterByBookId(id ...int64) NftRequestsQ
	FilterByNftId(id ...int64) NftRequestsQ
	FilterByCollectionAddress(address ...string) NftRequestsQ
	FilterByPayerAddress(address ...string) NftRequestsQ

	FilterUpdateById(id ...int64) NftRequestsQ
	UpdateStatus(newStatus resources.NftRequestStatus) NftRequestsQ
	UpdateLastUpdated(time time.Time) NftRequestsQ
	Update() error
}