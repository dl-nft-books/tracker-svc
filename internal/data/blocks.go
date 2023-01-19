package data

type Blocks struct {
	Id                 int64  `db:"id" structs:"-" json:"-"`
	ContractId         int64  `db:"contract_id" structs:"contract_id" json:"contract_id"`
	TransferBlock      uint64 `db:"transfer_block" structs:"transfer_block" json:"transfer_block"`
	UpdateBlock        uint64 `db:"update_block" structs:"update_block" json:"update_block"`
	VoucherUpdateBlock uint64 `db:"voucher_update_block" structs:"voucher_update_block" json:"voucher_update_block"`
}

type BlocksQ interface {
	New() BlocksQ

	FilterById(id ...int64) BlocksQ
	FilterByContractId(contractId ...int64) BlocksQ

	Insert(block Blocks) (int64, error)
	Get() (*Blocks, error)
	Select() ([]Blocks, error)

	UpdateTransferBlock(newTransferBlock uint64, id int64) error
	UpdateParamsUpdateBlock(newUpdateBlock uint64, id int64) error
	UpdateParamsVoucherUpdateBlock(newVoucherUpdateBlock uint64, id int64) error

	Delete(id int64) error
}
