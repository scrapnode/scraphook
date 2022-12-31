package cmd

import (
	corecmd "github.com/scrapnode/scrapcore/cmd"
	coredb "github.com/scrapnode/scrapcore/database/sql"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	command := &cobra.Command{
		Use:       "webhook",
		Short:     "webhook service commands",
		Example:   "scraphook webhook serve",
		ValidArgs: []string{"serve"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := corecmd.ChainPreRunE()(cmd, args); err != nil {
				return err
			}

			ctx := cmd.Context()
			provider := xconfig.FromContext(ctx)
			cfg, err := configs.New(provider)
			if err != nil {
				return err
			}
			ctx = configs.WithContext(ctx, cfg)

			logger := xlogger.New(cfg.Debug()).With("service", "scraphook.webhook")
			ctx = xlogger.WithContext(ctx, logger)

			if ok := corecmd.MustGetFlagBool(cmd, "auto-migrate"); ok {
				db, err := coredb.New(xlogger.WithContext(ctx, logger.With("fn", "cli.auto-migrate")), cfg.Database)
				if err != nil {
					logger.Fatal(err)
				}
				defer func() {
					if err := db.Disconnect(ctx); err != nil {
						logger.Error(err)
					}
				}()
				if err := db.Connect(ctx); err != nil {
					logger.Fatal(err)
				}

				if err := db.Migrate(ctx); err != nil {
					logger.Fatal(err)
				}
			}

			cmd.SetContext(ctx)
			return nil
		},
	}

	command.AddCommand(NewServe())

	command.PersistentFlags().BoolP("auto-migrate", "", false, "run migrate up automatically")
	return command
}
