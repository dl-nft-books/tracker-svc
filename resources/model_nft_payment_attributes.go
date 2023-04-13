/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftPaymentAttributes struct {
	// Book id
	BookId int64 `json:"book_id"`
	// Url to see the book
	BookUrl *string `json:"book_url,omitempty"`
	// Address of a book
	ContractAddress string `json:"contract_address"`
	// Floor nft price
	FloorPrice string `json:"floor_price"`
	// Price of a minted marketplace in $
	MintedTokenPrice string `json:"minted_token_price"`
	// Address of a nft collection
	NftAddress string `json:"nft_address"`
	// Nft id
	NftId int64 `json:"nft_id"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Timestamp when the user have purchased a book
	PurchaseTimestamp string `json:"purchase_timestamp"`
	// Token id
	TokenId int64 `json:"token_id"`
}
