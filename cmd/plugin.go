package cmd

import (
	"github.com/aunem/transpose/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pluginCmd)
}

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "manage plugins",
	Long:  utils.GetArt(),
}
