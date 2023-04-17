package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/dl-nft-books/tracker-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	paymentsTable = "payments"

	paymentsId              = "id"
	paymentsChainId         = "chain_id"
	paymentsTokenId         = "token_id"
	paymentsBookId          = "book_id"
	paymentsContractAddress = "contract_address"
	paymentsPayerAddress    = "payer_address"
	paymentsTokenAddress    = "token_address"
	paymentsAmount          = "amount"
	paymentsPrice           = "price"
	paymentsType            = "type"
)

type paymentsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewPaymentsQ(database *pgdb.DB) data.PaymentsQ {
	return &paymentsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", paymentsTable)).From(paymentsTable),
	}
}

func (q *paymentsQ) New() data.PaymentsQ {
	return NewPaymentsQ(q.database.Clone())
}

func (q *paymentsQ) Page(page pgdb.OffsetPageParams, cols ...string) data.PaymentsQ {
	if len(cols) == 0 {
		cols = append(cols, "id")
	}
	q.selector = page.ApplyTo(q.selector, cols...)
	return q
}

func (q *paymentsQ) Sort(sort pgdb.Sorts) data.PaymentsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *paymentsQ) FilterById(id ...int64) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsId: id})
	return q
}

func (q *paymentsQ) FilterByPayer(payer ...string) data.PaymentsQ {
	for i := range payer {
		payer[i] = strings.ToLower(payer[i])
	}
	q.selector = q.selector.Where(squirrel.Eq{fmt.Sprintf("lower(%v)", paymentsPayerAddress): payer})
	return q
}

func (q *paymentsQ) FilterByTokenAddress(tokenAddress ...string) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsTokenAddress: tokenAddress})
	return q
}

func (q *paymentsQ) FilterByChainId(chainId ...int64) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsChainId: chainId})
	return q
}

func (q *paymentsQ) FilterByTokenId(tokenId ...int64) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsTokenId: tokenId})
	return q
}

func (q *paymentsQ) FilterByBookId(bookId ...int64) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsBookId: bookId})
	return q
}

func (q *paymentsQ) FilterByType(paymentType ...int8) data.PaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{paymentsType: paymentType})
	return q
}

func (q *paymentsQ) FilterByContractAddress(contractAddress ...string) data.PaymentsQ {
	for i := range contractAddress {
		contractAddress[i] = strings.ToLower(contractAddress[i])
	}
	q.selector = q.selector.Where(squirrel.Eq{fmt.Sprintf("lower(%v)", paymentsContractAddress): contractAddress})
	return q
}

func (q *paymentsQ) Select() (payments []data.Payment, err error) {
	err = q.database.Select(&payments, q.selector)
	return
}

func (q *paymentsQ) Get() (*data.Payment, error) {
	var payment data.Payment
	err := q.database.Get(&payment, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &payment, err
}

func (q *paymentsQ) Insert(payment data.Payment) (id int64, err error) {
	statement := squirrel.Insert(paymentsTable).
		Suffix("returning id").
		SetMap(structs.Map(&payment))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *paymentsQ) Delete(id int64) error {
	statement := squirrel.Delete(paymentsTable).Where(squirrel.Eq{paymentsId: id})
	return q.database.Exec(statement)
}
