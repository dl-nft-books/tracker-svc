/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PaymentAttributes struct {
	// Amount of tokens paid
	Amount string `json:"amount"`
	// Address of a user who bought a book
	PayerAddress string `json:"payer_address"`
	// Price of a token in $
	Price string `json:"price"`
	// Token metadata information
	Token Token `json:"token"`
}
