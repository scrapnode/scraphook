package cmd

import (
	corecmd "github.com/scrapnode/scrapcore/cmd"
	webhookcmd "github.com/scrapnode/scraphook/webhook/cmd"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	command := corecmd.New()

	command.AddCommand(webhookcmd.New())
	return command
}
