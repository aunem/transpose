package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aunem/transpose/config"
	"github.com/aunem/transpose/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string
var namespace string
var conf *config.Transpose

var rootCmd = &cobra.Command{
	Use:   "transpose",
	Short: "Transpose is a lightweight composable proxy",
	Long:  utils.GetArt(),
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&name, "config-name", "", "name of the kubernetes config to sync to")
	rootCmd.PersistentFlags().StringVar(&namespace, "config-namespace", "local", "namespace of the kubernetes config to sync to, use 'local' for a local config")
	viper.BindPFlag("config-name", rootCmd.PersistentFlags().Lookup("config-name"))
	viper.BindPFlag("config-namespace", rootCmd.PersistentFlags().Lookup("config-namespace"))
}

func initConfig() {
	r := strings.NewReplacer("_", "-")
	viper.SetEnvKeyReplacer(r)
	viper.AutomaticEnv()
	var err error
	name := viper.GetString("config-name")
	namespace := viper.GetString("config-namespace")
	conf, err = config.LoadConfig(name, namespace)
	if err != nil {
		log.Fatalf("could not load config: %+v", err)
	}
	log.Infof("loaded config: %+v", conf)
}

// Execute is the root cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
