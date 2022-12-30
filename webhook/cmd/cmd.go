package cmd

import (
	corecmd "github.com/scrapnode/scrapcore/cmd"
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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
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

			cmd.SetContext(ctx)
			return nil
		},
	}

	command.AddCommand(NewServe())
	return command
}
