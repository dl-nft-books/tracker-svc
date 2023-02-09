package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type NftPayment struct {
	Id                int64     `db:"id" structs:"-" json:"-"`
	ContractId        int64     `db:"contract_id" structs:"contract_id" json:"contract_id"`
	ContractAddress   string    `db:"contract_address" structs:"contract_address"`
	PayerAddress      string    `db:"payer_address" structs:"payer_address"`
	NftAddress        string    `db:"nft_address" structs:"nft_address"`
	NftId             int64     `db:"nft_id" structs:"nft_id" json:"nft_id"`
	FloorPrice        string    `db:"floor_price" structs:"floor_price" json:"floor_price"`
	PriceMinted       string    `db:"price_minted" structs:"price_minted"`
	BookUrl           string    `db:"book_url" structs:"book_url"`
	PurchaseTimestamp time.Time `db:"purchase_timestamp" structs:"purchase_timestamp"`
}

type NftPaymentsQ interface {
	New() NftPaymentsQ

	FilterById(id ...int64) NftPaymentsQ
	FilterByPayer(payer ...string) NftPaymentsQ
	FilterByNftAddress(nftAddress ...string) NftPaymentsQ
	FilterByContractAddress(contractAddress ...string) NftPaymentsQ
	FilterByContractId(contractId ...int64) NftPaymentsQ
	FilterByBookUrl(bookUrl ...string) NftPaymentsQ

	Get() (*NftPayment, error)
	Select() ([]NftPayment, error)

	Sort(sort pgdb.Sorts) NftPaymentsQ
	Page(page pgdb.OffsetPageParams) NftPaymentsQ

	Insert(payment NftPayment) (id int64, err error)
	Delete(id int64) error
}

//func (p *NftPayment) Resource() (*resources.Payment, error) {
//	return &resources.Payment{
//		Key: resources.NewKeyInt64(p.Id, resources.PAYMENT),
//		Attributes: resources.PaymentAttributes{
//			Amount:            p.Amount,
//			PayerAddress:      p.PayerAddress,
//			PaymentTokenPrice: p.PriceToken,
//			MintedTokenPrice:  p.PriceMinted,
//			PurchaseTimestamp: p.PurchaseTimestamp.Format(timestampFormat),
//			BookUrl:           p.BookUrl,
//			Erc20Data: resources.Erc20Data{
//				Address:  p.TokenAddress,
//				Name:     p.TokenName,
//				Symbol:   p.TokenSymbol,
//				Decimals: int32(p.TokenDecimals),
//			},
//		},
//	}, nil
//}
