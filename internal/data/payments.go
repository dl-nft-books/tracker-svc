package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

type Payment struct {
	Id              int64  `db:"id" structs:"-" json:"-"`
	ContractId      int64  `db:"contract_id" structs:"contract_id" json:"contract_id"`
	ContractAddress string `db:"contract_address" structs:"contract_address"`
	PayerAddress    string `db:"payer_address" structs:"payer_address"`
	TokenAddress    string `db:"token_address" structs:"token_address"`
	TokenSymbol     string `db:"token_symbol" structs:"token_symbol"`
	TokenName       string `db:"token_name" structs:"token_name"`
	Amount          string `db:"amount" structs:"amount"`
	Price           string `db:"price" structs:"price"`
}

type PaymentsQ interface {
	New() PaymentsQ

	FilterById(id ...int64) PaymentsQ
	FilterByPayer(payer ...string) PaymentsQ
	FilterByTokenAddress(tokenAddress ...string) PaymentsQ
	FilterByContractAddress(contractAddress ...string) PaymentsQ
	FilterByContractId(contractId ...int64) PaymentsQ

	Get() (*Payment, error)
	Select() ([]Payment, error)

	Sort(sort pgdb.Sorts) PaymentsQ
	Page(page pgdb.OffsetPageParams) PaymentsQ

	Insert(payment Payment) (id int64, err error)
	Delete(id int64) error
}

func (p *Payment) Resource() resources.Payment {
	return resources.Payment{
		Key: resources.NewKeyInt64(p.Id, resources.PAYMENT),
		Attributes: resources.PaymentAttributes{
			Amount:       p.Amount,
			PayerAddress: p.PayerAddress,
			Price:        p.Price,
			Token: resources.Token{
				Address: p.TokenAddress,
				Name:    p.TokenName,
				Symbol:  p.TokenSymbol,
			},
		},
	}
}
