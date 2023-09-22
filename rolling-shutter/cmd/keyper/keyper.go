package keyper

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/shversion"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/keyper/database"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/configuration/command"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/db"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/service"
)

func Cmd() *cobra.Command {
	builder := command.Build(
		main,
		command.Usage(
			"Run a Shutter keyper node",
			`This command runs a keyper node. It will connect to both an Ethereum and a
Shuttermint node which have to be started separately in advance.`,
		),
		command.WithGenerateConfigSubcommand(),
	)
	builder.AddInitDBCommand(initDB)
	return builder.Command()
}

func main(config *keyper.Config) error {
	log.Info().
		Str("version", shversion.Version()).
		Str("address", config.GetAddress().Hex()).
		Str("shuttermint", config.Shuttermint.ShuttermintURL).
		Msg("starting keyper")

	return service.RunWithSighandler(context.Background(), keyper.New(config))
}

func initDB(cfg *keyper.Config) error {
	ctx := context.Background()

	dbpool, err := pgxpool.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to database")
	}
	defer dbpool.Close()

	return db.InitDB(ctx, dbpool, database.Definition.Name(), database.Definition)
}
