package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/aunem/transpose/pkg/listener"
	"github.com/aunem/transpose/utils"
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
	log.Infof("starting server...")
	m, err := listener.NewManager(conf.Spec)
	if err != nil {
		log.Fatalf("problem creating listener manager: %+v", err)
	}
	err = m.ExecListener()
	if err != nil {
		log.Fatalf("problem executing listener: %+v", err)
	}
}
