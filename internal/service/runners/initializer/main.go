package initializer

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners/combiners"
)

type RunnerInitializer struct {
	database pgdb.DB
	combiner combiners.TokenCombiner
}

func NewRunnerInitializer(cfg config.Config, ctx context.Context) {

}
