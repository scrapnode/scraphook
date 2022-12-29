package cmd

import (
	corecmd "github.com/scrapnode/scrapcore/pkg/cmd"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	command := corecmd.New()

	return command
}
