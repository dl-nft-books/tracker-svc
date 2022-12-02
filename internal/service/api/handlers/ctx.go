package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	booker "gitlab.com/tokend/nft-books/book-svc/connector/api"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/data"
	generatorer "gitlab.com/tokend/nft-books/generator-svc/connector/api"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
	bookerCtxKey
	generatorerCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDB(db data.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, db)
	}
}

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbCtxKey).(data.DB).New()
}

func CtxBooker(entry booker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, bookerCtxKey, entry)
	}
}

func Booker(r *http.Request) booker.Connector {
	return r.Context().Value(bookerCtxKey).(booker.Connector)
}

func CtxGeneratorer(entry generatorer.Connector) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, generatorerCtxKey, entry)
	}
}

func Generatorer(r *http.Request) generatorer.Connector {
	return r.Context().Value(generatorerCtxKey).(generatorer.Connector)
}
