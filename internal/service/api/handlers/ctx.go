package handlers

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data/external"
	"net/http"

	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	booksQCtxKey
	trackerDBCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxBooksQ(q external.BookQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, booksQCtxKey, q)
	}
}

func CtxTrackerDB(db data.TrackerDB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, trackerDBCtxKey, db)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BooksQ(r *http.Request) external.BookQ {
	return r.Context().Value(booksQCtxKey).(external.BookQ).New()
}

func TrackerDB(r *http.Request) data.TrackerDB {
	return r.Context().Value(trackerDBCtxKey).(data.TrackerDB).New()
}
