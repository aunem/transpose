package cmd

import (
	"os"

	"github.com/aunem/transpose/pkg/utils"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove all plugin objects from bin directory",
	Long:  utils.GetArt(),
	Run:   Clean,
}

func init() {
	pluginCmd.AddCommand(cleanCmd)
}

// Clean plugins
func Clean(cmd *cobra.Command, args []string) {
	err := os.RemoveAll("./bin")
	if err != nil {
		log.Fatal(err)
	}
	utils.MakeBins()
	log.Info("sucessfully cleaned plugins")
}
