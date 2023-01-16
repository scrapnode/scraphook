package cmd

import (
	"context"
	corecmd "github.com/scrapnode/scrapcore/cmd"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/attempt/services"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewServe() *cobra.Command {
	command := &cobra.Command{
		Use:       "serve",
		Example:   `scraphook attempt serve webserver`,
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"trigger"},
		PreRunE:   corecmd.ChainPreRunE(),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			logger := xlogger.FromContext(ctx).
				With("fn", "cli.serve")

			name := args[0]
			srv, err := services.New(ctx, name)
			if err != nil {
				logger.Fatal(err)
			}

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				if err := srv.Start(ctx); err != nil {
					logger.Error(err)
					cancel()
					return
				}
				if err := srv.Run(ctx); err != nil {
					logger.Error(err)
					cancel()
					return
				}
			}()

			logger.Infow("running", "service_name", name)
			// Listen for the interrupt signal.
			<-ctx.Done()
			// make sure once we stop process, we cancel all the execution
			cancel()

			logger.Info("shutting down gracefully, press Ctrl+C again to force")
			// The context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			// Because the context channel is done, so we could not reuse it, we have to use original context here
			ctx, cancel = context.WithTimeout(cmd.Context(), 11*time.Second)
			go func() {
				if err := srv.Stop(ctx); err != nil {
					logger.Error(err)
				}
				cancel()
			}()
			<-ctx.Done()
		},
	}

	return command
}
