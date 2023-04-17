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
	nftPaymentsTable = "nft_payments"

	nftPaymentsId              = "id"
	nftPaymentsChainId         = "chain_id"
	nftPaymentsContractAddress = "contract_address"
	nftPaymentsPayerAddress    = "payer_address"
	nftPaymentsNftAddress      = "nft_address"
	nftPaymentsNftId           = "nft_id"
	nftPaymentsBookUrl         = "book_url"
)

type nftPaymentsQ struct {
	database *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewNftPaymentsQ(database *pgdb.DB) data.NftPaymentsQ {
	return &nftPaymentsQ{
		database: database,
		selector: squirrel.Select(fmt.Sprintf("%s.*", nftPaymentsTable)).From(nftPaymentsTable),
	}
}

func (q *nftPaymentsQ) New() data.NftPaymentsQ {
	return NewNftPaymentsQ(q.database.Clone())
}

func (q *nftPaymentsQ) Page(page pgdb.OffsetPageParams, cols ...string) data.NftPaymentsQ {
	if len(cols) == 0 {
		cols = append(cols, "id")
	}
	q.selector = page.ApplyTo(q.selector, cols...)
	return q
}

func (q *nftPaymentsQ) Sort(sort pgdb.Sorts) data.NftPaymentsQ {
	q.selector = sort.ApplyTo(q.selector, map[string]string{
		"id": "id",
	})

	return q
}

func (q *nftPaymentsQ) FilterById(id ...int64) data.NftPaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftPaymentsId: id})
	return q
}

func (q *nftPaymentsQ) FilterByPayer(payer ...string) data.NftPaymentsQ {
	for i := range payer {
		payer[i] = strings.ToLower(payer[i])
	}
	q.selector = q.selector.Where(squirrel.Eq{fmt.Sprintf("lower(%v)", nftPaymentsPayerAddress): payer})
	return q
}

func (q *nftPaymentsQ) FilterByNftAddress(nftAddress ...string) data.NftPaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftPaymentsNftAddress: nftAddress})
	return q
}

func (q *nftPaymentsQ) FilterByNftId(nftId ...int64) data.NftPaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftPaymentsNftId: nftId})
	return q
}

func (q *nftPaymentsQ) FilterByBookUrl(bookUrl ...string) data.NftPaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftPaymentsBookUrl: bookUrl})
	return q
}

func (q *nftPaymentsQ) FilterByChainId(chainId ...int64) data.NftPaymentsQ {
	q.selector = q.selector.Where(squirrel.Eq{nftPaymentsChainId: chainId})
	return q
}

func (q *nftPaymentsQ) FilterByContractAddress(contractAddress ...string) data.NftPaymentsQ {
	for i := range contractAddress {
		contractAddress[i] = strings.ToLower(contractAddress[i])
	}
	q.selector = q.selector.Where(squirrel.Eq{fmt.Sprintf("lower(%v)", nftPaymentsContractAddress): contractAddress})
	return q
}

func (q *nftPaymentsQ) OrderBy(column ...string) data.NftPaymentsQ {

	q.selector = q.selector.OrderBy()
	return q
}

func (q *nftPaymentsQ) Select() (nftPayments []data.NftPayment, err error) {
	err = q.database.Select(&nftPayments, q.selector)
	return
}

func (q *nftPaymentsQ) Get() (*data.NftPayment, error) {
	var payment data.NftPayment
	err := q.database.Get(&payment, q.selector)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &payment, err
}

func (q *nftPaymentsQ) Insert(payment data.NftPayment) (id int64, err error) {
	statement := squirrel.Insert(nftPaymentsTable).
		Suffix("returning id").
		SetMap(structs.Map(&payment))

	err = q.database.Get(&id, statement)
	return id, err
}

func (q *nftPaymentsQ) Delete(id int64) error {
	statement := squirrel.Delete(nftPaymentsTable).Where(squirrel.Eq{nftPaymentsId: id})
	return q.database.Exec(statement)
}
