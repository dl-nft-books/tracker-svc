package external

import (
	"time"

	"gitlab.com/tokend/nft-books/contract-tracker/resources"
)

type BookQ interface {
	New() BookQ
	Get() (*Book, error)
	Select() ([]Book, error)
	FilterActual() BookQ
	FilterByID(id ...int64) BookQ
	FilterByContractAddress(address ...string) BookQ
}

type Book struct {
	ID              int64                  `db:"id" structs:"-"`
	Title           string                 `db:"title" structs:"title"`
	Description     string                 `db:"description" structs:"description"`
	CreatedAt       time.Time              `db:"created_at" structs:"created_at"`
	Price           string                 `db:"price" structs:"price"`
	ContractAddress string                 `db:"contract_address" structs:"contract_address"`
	ContractName    string                 `db:"contract_name" structs:"contract_name"`
	ContractSymbol  string                 `db:"contract_symbol" structs:"contract_symbol"`
	ContractVersion string                 `db:"contract_version" structs:"contract_version"`
	Banner          string                 `db:"banner" structs:"banner"`
	File            string                 `db:"file" structs:"file"`
	Deleted         bool                   `db:"deleted" structs:"-"`
	ChainID         int64                  `db:"chain_id" structs:"chain_id"`
	TokenId         int64                  `db:"token_id" structs:"token_id"`
	DeployStatus    resources.DeployStatus `db:"deploy_status" structs:"deploy_status"`
	LastBlock       uint64                 `db:"last_block" structs:"last_block"`
}
