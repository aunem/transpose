package cmd

import (
	"os"

	"github.com/aunem/transpose/config"
	"github.com/aunem/transpose/pkg/listener"
	"github.com/aunem/transpose/pkg/middleware"
	"github.com/aunem/transpose/pkg/roundtrip"
	"github.com/aunem/transpose/pkg/utils"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var localBuild bool
var build bool

var reolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "resolve plugins",
	Long:  utils.GetArt(),
	Run:   Resolve,
}

func init() {
	pluginCmd.AddCommand(reolveCmd)
	initCmd.Flags().BoolVarP(&localBuild, "local", "l", false, "build plugin from local gopath")
	initCmd.Flags().BoolVarP(&build, "build", "b", false, "resolve all plugins anew")
}

// Resolve plugins
func Resolve(cmd *cobra.Command, args []string) {
	if build {
		err := os.RemoveAll("./bin")
		if err != nil {
			log.Fatal(err)
		}
	}
	utils.MakeBins()
	log.Info("loading config...")
	c, err := config.LoadConfig("", "local")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving middleware plugin...")
	mw, err := middleware.NewManager(c.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving roundtrip plugin...")
	rt, err := roundtrip.NewManager(c.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving listener plugin...")
	_, err = listener.NewManager(c.Spec, mw, rt)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("all plugins resolved, objects can be found in ./bin folder")
}
