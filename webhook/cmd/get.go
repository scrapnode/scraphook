package cmd

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	corecmd "github.com/scrapnode/scrapcore/cmd"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/spf13/cobra"
)

func NewGet() *cobra.Command {
	command := &cobra.Command{
		Use:       "get",
		Example:   `scraphook webhook get configs`,
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"configs"},
		PreRunE:   corecmd.ChainPreRunE(),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			info := args[0]
			if info == "configs" {
				cfg := configs.FromContext(ctx)
				bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalIndent(cfg, "", "  ")
				if err != nil {
					panic(err)
				}

				fmt.Println(string(bytes))
			}
		},
	}

	return command
}
