package data

type Blocks struct {
	Id              int64  `db:"id" structs:"-" json:"-"`
	ContractAddress string `db:"contract_address" structs:"contract_address"`
	ChainId         int64  `db:"chain_id" structs:"chain_id"`
	MintBlock       uint64 `db:"mint_block" structs:"mint_block" json:"mint_block"`
	MintByNFTBlock  uint64 `db:"mint_by_nft_block" structs:"mint_by_nft_block" json:"mint_by_nft_block"`
}

type BlocksQ interface {
	New() BlocksQ

	FilterById(id ...int64) BlocksQ
	FilterByChainId(id ...int64) BlocksQ
	FilterByContractAddress(contractAddress ...string) BlocksQ

	Insert(block Blocks) (int64, error)
	Get() (*Blocks, error)
	Select() ([]Blocks, error)

	UpdateMintBlock(newMintBlock uint64, chainId int64) error
	UpdateMintByNFTBlock(newMintByNFTBlock uint64, chainId int64) error

	Delete(chainId int64) error
}
