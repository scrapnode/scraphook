package cmd

import (
	"context"
	corecmd "github.com/scrapnode/scrapcore/cmd"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	command := &cobra.Command{
		Use:       "webhook",
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

			logger := xlogger.New(cfg.Debug()).
				With("service_group", "scraphook.webhook").
				With("version", cfg.Version)
			ctx = xlogger.WithContext(ctx, logger)

			if err := runDBTasks(cmd, ctx); err != nil {
				return err
			}

			cmd.SetContext(ctx)
			return nil
		},
	}

	command.AddCommand(NewServe())
	command.AddCommand(NewGet())

	command.PersistentFlags().BoolP("auto-migrate", "", false, "run migrate up automatically")
	command.PersistentFlags().StringArrayP("seeds", "", []string{}, "seed files that will be run before start your application")

	return command
}

func runDBTasks(cmd *cobra.Command, ctx context.Context) error {
	shouldMigrate := corecmd.MustGetFlagBool(cmd, "auto-migrate")
	seeds := corecmd.MustGetFlagStringArray(cmd, "seeds")
	if !shouldMigrate && len(seeds) == 0 {
		return nil
	}

	cfg := configs.FromContext(ctx)
	logger := xlogger.FromContext(ctx).With("fn", "cli.auto-migrate")
	ctx = xlogger.WithContext(ctx, logger)

	db, err := database.New(ctx, cfg.Database)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			logger.Error(err)
		}
	}()
	if err := db.Connect(ctx); err != nil {
		return err
	}

	if shouldMigrate {
		if err := db.Migrate(ctx); err != nil {
			return err
		}
	}

	if len(seeds) > 0 {
		if err := db.Seed(ctx, seeds); err != nil {
			return err
		}
	}

	return nil
}
