package cmd

import (
	"context"
	corecmd "github.com/scrapnode/scrapcore/cmd"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
	"github.com/scrapnode/scraphook/webhook/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewServe() *cobra.Command {
	command := &cobra.Command{
		Use:     "serve",
		Short:   "serve webhook servers",
		Example: "scraphook webhook serve",
		PreRunE: corecmd.ChainPreRunE(),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			cfg := configs.FromContext(ctx)

			infra, err := infrastructure.New(ctx, cfg)
			if err != nil {
				infra.Logger.Fatal("could not initial infrastructure", "error", err.Error())
			}

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			srv := server.New(ctx, infra)

			go func() {
				if err := srv.Start(ctx); err != nil {
					infra.Logger.Error(err)
					cancel()
					return
				}
				if err := srv.Run(); err != nil {
					infra.Logger.Error(err)
					cancel()
					return
				}
			}()

			infra.Logger.Info("running")
			// Listen for the interrupt signal.
			<-ctx.Done()
			// make sure once we stop process, we cancel all the execution
			cancel()

			infra.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
			// The context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			// Because the context channel is done, so we could not reuse it, we have to use original context here
			ctx, cancel = context.WithTimeout(cmd.Context(), 11*time.Second)
			go func() {
				if err = srv.Stop(ctx); err != nil {
					infra.Logger.Error(err)
				}
				cancel()
			}()
			<-ctx.Done()
		},
	}

	return command
}
