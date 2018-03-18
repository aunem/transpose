package cmd

import (
	"fmt"

	"github.com/aunem/transpose/utils"
	"github.com/spf13/cobra"
)

var ptyp string
var pname string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a plugin repository",
	Long:  utils.GetArt(),
	Run:   Init,
}

func init() {
	pluginCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&ptyp, "type", "t", "", "type of plugin (listener, middleware, or roundtrip)")
	initCmd.Flags().StringVarP(&pname, "name", "n", "", "name of plugin")
	initCmd.MarkFlagRequired("type")
	initCmd.MarkFlagRequired("name")
}

// Init plugins
func Init(cmd *cobra.Command, args []string) {
	fmt.Println("not yet implemented")
}
