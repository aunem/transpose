package cmd

import (
	"fmt"

	res "github.com/aunem/transpose/pkg/resolve"
	"github.com/aunem/transpose/pkg/template"
	"github.com/aunem/transpose/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ptyp, pname, ppackage string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a plugin repository",
	Long:  utils.GetArt(),
	Run:   Init,
}

func init() {
	pluginCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&ptyp, "type", "t", "", "type of plugin (listener, middleware, or roundtrip) (required)")
	initCmd.Flags().StringVarP(&pname, "name", "n", "", "name of plugin (required)")
	initCmd.Flags().StringVarP(&ppackage, "package", "p", "", "golang package (required)")
	initCmd.MarkFlagRequired("type")
	initCmd.MarkFlagRequired("name")
	initCmd.MarkFlagRequired("package")
}

// Init plugins
func Init(cmd *cobra.Command, args []string) {

	fmt.Println("creating templates...")
	p := template.NewPlugin(pname, ppackage, ptyp)
	err := p.Template()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resolving dependencies...")
	res.Init()
	fmt.Println("successfully templated plugin repo!")
}
