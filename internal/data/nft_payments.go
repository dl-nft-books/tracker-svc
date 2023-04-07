package data

import (
	"github.com/dl-nft-books/tracker-svc/resources"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type NftPayment struct {
	Id                int64     `db:"id" structs:"-" json:"-"`
	ContractAddress   string    `db:"contract_address" structs:"contract_address"`
	PayerAddress      string    `db:"payer_address" structs:"payer_address"`
	NftAddress        string    `db:"nft_address" structs:"nft_address"`
	NftId             int64     `db:"nft_id" structs:"nft_id" json:"nft_id"`
	FloorPrice        string    `db:"floor_price" structs:"floor_price" json:"floor_price"`
	PriceMinted       string    `db:"price_minted" structs:"price_minted"`
	BookUrl           string    `db:"book_url" structs:"book_url"`
	ChainId           int64     `db:"chain_id" structs:"chain_id"`
	PurchaseTimestamp time.Time `db:"purchase_timestamp" structs:"purchase_timestamp"`
}

type NftPaymentsQ interface {
	New() NftPaymentsQ

	FilterById(id ...int64) NftPaymentsQ
	FilterByPayer(payer ...string) NftPaymentsQ
	FilterByNftAddress(nftAddress ...string) NftPaymentsQ
	FilterByNftId(nftId ...int64) NftPaymentsQ
	FilterByContractAddress(contractAddress ...string) NftPaymentsQ
	FilterByChainId(chainId ...int64) NftPaymentsQ
	FilterByBookUrl(bookUrl ...string) NftPaymentsQ
	OrderBy(column ...string) NftPaymentsQ

	Get() (*NftPayment, error)
	Select() ([]NftPayment, error)

	Sort(sort pgdb.Sorts) NftPaymentsQ
	Page(page pgdb.OffsetPageParams, cols ...string) NftPaymentsQ

	Insert(payment NftPayment) (id int64, err error)
	Delete(id int64) error
}

func (p *NftPayment) Resource() (*resources.NftPayment, error) {
	return &resources.NftPayment{
		Key: resources.NewKeyInt64(p.Id, resources.PAYMENT),
		Attributes: resources.NftPaymentAttributes{
			NftAddress:        p.NftAddress,
			NftId:             int32(p.NftId),
			FloorPrice:        p.FloorPrice,
			PayerAddress:      p.PayerAddress,
			MintedTokenPrice:  p.PriceMinted,
			PurchaseTimestamp: p.PurchaseTimestamp.Format(timestampFormat),
			BookUrl:           p.BookUrl,
		},
	}, nil
}
