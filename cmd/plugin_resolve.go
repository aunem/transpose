package cmd

import (
	"fmt"

	"github.com/aunem/transpose/utils"
	"github.com/spf13/cobra"
)

var reolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "resolve plugins",
	Long:  utils.GetArt(),
	Run:   Resolve,
}

func init() {
	pluginCmd.AddCommand(reolveCmd)
}

// Resolve plugins
func Resolve(cmd *cobra.Command, args []string) {
	fmt.Println("not yet implemented")
}
