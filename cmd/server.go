package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/aunem/transpose/pkg/listener"
	"github.com/aunem/transpose/pkg/middleware"
	"github.com/aunem/transpose/pkg/roundtrip"
	"github.com/aunem/transpose/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts a new transpose server",
	Long:  utils.GetArt(),
	Run:   Serve,
}

// Serve starts a new transpose server
func Serve(cmd *cobra.Command, args []string) {
	InitConfig()
	log.Info("finding or creating bins...")
	utils.MakeBins()
	log.Info("resolving middleware plugin...")
	mw, err := middleware.NewManager(conf.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving roundtrip plugin...")
	rt, err := roundtrip.NewManager(conf.Spec)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("resolving listener plugin")
	m, err := listener.NewManager(conf.Spec, mw, rt)
	if err != nil {
		log.Fatalf("problem creating listener manager: %+v", err)
	}
	log.Infof("starting listener...")
	err = m.ExecListener()
	if err != nil {
		log.Fatalf("problem executing listener: %+v", err)
	}
}
