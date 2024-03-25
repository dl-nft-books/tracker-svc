/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PaymentAttributes struct {
	// Amount of tokens paid
	Amount string `json:"amount"`
	// Url to see the book's banner
	BannerUrl string `json:"banner_url"`
	// Book id
	BookId int64 `json:"book_id"`
	// Address of a book
	ContractAddress string `json:"contract_address"`
	// Token metadata information
	Erc20Data Erc20Data `json:"erc20_data"`
	// Price of a minted marketplace in $
	MintedTokenPrice string `json:"minted_token_price"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Price of a payment token (which the minted token was bought by) in $
	PaymentTokenPrice string `json:"payment_token_price"`
	// Timestamp when the user have purchased a book
	PurchaseTimestamp string `json:"purchase_timestamp"`
	// Token id
	TokenId int64 `json:"token_id"`
	// Event type
	Type string `json:"type"`
}
