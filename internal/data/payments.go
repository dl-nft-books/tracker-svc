package data

import (
	"time"

	"github.com/dl-nft-books/tracker-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const timestampFormat = "2006-01-02"

type Payment struct {
	Id                int64     `db:"id" structs:"-" json:"-"`
	ContractAddress   string    `db:"contract_address" structs:"contract_address"`
	PayerAddress      string    `db:"payer_address" structs:"payer_address"`
	TokenAddress      string    `db:"token_address" structs:"token_address"`
	TokenSymbol       string    `db:"token_symbol" structs:"token_symbol"`
	TokenName         string    `db:"token_name" structs:"token_name"`
	TokenDecimals     uint8     `db:"token_decimals" structs:"token_decimals"`
	Amount            string    `db:"amount" structs:"amount"`
	PriceToken        string    `db:"price_token" structs:"price_token"`
	PriceMinted       string    `db:"price_minted" structs:"price_minted"`
	BannerLink        string    `db:"banner_link" structs:"banner_link"`
	ChainId           int64     `db:"chain_id" structs:"chain_id"`
	NftId             int64     `db:"nft_id" structs:"nft_id"`
	TokenId           int64     `db:"token_id" structs:"token_id"`
	BookId            int64     `db:"book_id" structs:"book_id"`
	PurchaseTimestamp time.Time `db:"purchase_timestamp" structs:"purchase_timestamp"`
	Type              int8      `db:"type" structs:"type"`
}

type PaymentsQ interface {
	New() PaymentsQ

	FilterById(id ...int64) PaymentsQ
	FilterByPayer(payer ...string) PaymentsQ
	FilterByTokenAddress(tokenAddress ...string) PaymentsQ
	FilterByContractAddress(contractAddress ...string) PaymentsQ
	FilterByChainId(chainId ...int64) PaymentsQ
	FilterByTokenId(tokenId ...int64) PaymentsQ
	FilterByBookId(bookId ...int64) PaymentsQ
	FilterByType(paymentType ...int8) PaymentsQ

	Get() (*Payment, error)
	Select() ([]Payment, error)

	Sort(sort pgdb.Sorts) PaymentsQ
	Page(page pgdb.OffsetPageParams, cols ...string) PaymentsQ

	Insert(payment Payment) (id int64, err error)
	Delete(id int64) error
}

func (p *Payment) Resource() (*resources.Payment, error) {
	return &resources.Payment{
		Key: resources.NewKeyInt64(p.Id, resources.PAYMENT),
		Attributes: resources.PaymentAttributes{
			ContractAddress:   p.ContractAddress,
			TokenId:           p.TokenId,
			BookId:            p.BookId,
			Amount:            p.Amount,
			PayerAddress:      p.PayerAddress,
			PaymentTokenPrice: p.PriceToken,
			MintedTokenPrice:  p.PriceMinted,
			BannerUrl:         p.BannerLink,
			PurchaseTimestamp: p.PurchaseTimestamp.Format(timestampFormat),
			Type:              resources.TokenPurchasedEventType(p.Type).String(),
			Erc20Data: resources.Erc20Data{
				Address:  p.TokenAddress,
				Name:     p.TokenName,
				Symbol:   p.TokenSymbol,
				Decimals: int32(p.TokenDecimals),
			},
		},
	}, nil
}
