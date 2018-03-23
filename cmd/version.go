package cmd

import (
	"fmt"

	"github.com/aunem/transpose/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Transpose",
	Long:  utils.GetArt(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not yet implemented")
	},
}
