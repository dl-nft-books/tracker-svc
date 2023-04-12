package models

import "github.com/dl-nft-books/book-svc/resources"

type (
	GetBookResponse   resources.BookResponse
	ListBooksResponse resources.BookListResponse
	GetBookById struct {
		ChainId []int64 `filter:"chain_id"`
	}
)
