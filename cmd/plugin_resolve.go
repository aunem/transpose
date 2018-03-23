package cmd

import (
	"github.com/aunem/transpose/config"
	"github.com/aunem/transpose/pkg/listener"
	"github.com/aunem/transpose/pkg/middleware"
	"github.com/aunem/transpose/pkg/roundtrip"
	"github.com/aunem/transpose/utils"
	log "github.com/sirupsen/logrus"

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
	log.Info("loading config...")
	c, err := config.LoadConfig("", "local")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving listener plugin...")
	_, err = listener.NewManager(c.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving middleware plugin...")
	_, err = middleware.NewManager(c.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving roundtrip plugin...")
	_, err = roundtrip.NewManager(c.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("all plugins resolved, objects can be found in ./plugins folder")
}
