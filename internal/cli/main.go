package cli

import (
	"context"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/api"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/service/runners"

	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	// Cli commands
	var (
		app = kingpin.New("contract-tracker", "service for tracking contract events")

		runCmd      = app.Command("run", "run command")
		apiCmd      = runCmd.Command("api", "run api handlers")
		trackingCmd = runCmd.Command("tracking", "run events tracking")

		migrateCmd     = app.Command("migrate", "migrate command")
		migrateUpCmd   = migrateCmd.Command("up", "migrate db up")
		migrateDownCmd = migrateCmd.Command("down", "migrate db down")
	)

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case apiCmd.FullCommand():
		err = api.Run(cfg)
	case trackingCmd.FullCommand():
		err = runners.Run(cfg, context.Background())
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}

	return true
}
