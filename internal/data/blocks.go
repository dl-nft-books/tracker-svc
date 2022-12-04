package data

type Blocks struct {
	Id            int64  `db:"id" structs:"-" json:"-"`
	ContractId    int64  `db:"contract_id" structs:"contract_id" json:"contract_id"`
	TransferBlock uint64 `db:"transfer_block" structs:"transfer_block" json:"transfer_block"`
	UpdateBlock   uint64 `db:"update_block" structs:"update_block" json:"update_block"`
}

type BlocksQ interface {
	New() BlocksQ

	FilterById(id ...int64) BlocksQ
	FilterByContractId(contractId ...int64) BlocksQ

	Get() (*Blocks, error)
	Select() ([]Blocks, error)

	Delete(id int64) error
}
