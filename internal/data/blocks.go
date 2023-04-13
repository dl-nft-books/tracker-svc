package data

type Blocks struct {
	Id                  int64  `db:"id" structs:"-" json:"-"`
	ContractAddress     string `db:"contract_address" structs:"contract_address"`
	ChainId             int64  `db:"chain_id" structs:"chain_id"`
	TokenPurchasedBlock uint64 `db:"token_purchased_block" structs:"token_purchased_block" json:"token_purchased_block"`
}

type BlocksQ interface {
	New() BlocksQ

	FilterById(id ...int64) BlocksQ
	FilterByChainId(id ...int64) BlocksQ
	FilterByContractAddress(contractAddress ...string) BlocksQ

	Insert(block Blocks) (int64, error)
	Get() (*Blocks, error)
	Select() ([]Blocks, error)

	UpdateTokenPurchasedBlockColumn(newMintBlock uint64, chainId int64) error

	Delete(chainId int64) error
}
