package listener

var mainTpl = `package main

import (
	"fmt"
	"path"

	"github.com/aunem/transpose/pkg/context"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type {{ .Name }}Plugin struct {}

// ListenerPlugin exports the plugin struct
var ListenerPlugin {{ .Name }}Plugin

// Spec exports the spec data
var Spec {{ .Name }}Spec

func main() {}

// Listen implements the listener plugin inerface
func (h *{{ .Name }}Listener) Listen(mw *middleware.Manager, rt *roundtrip.Manager) error {
	// Implement listener code here
}

func (p *{{ .Name }}Plugin) Init(spec interface{}) error {
	b, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &Spec)
	if err != nil {
		return err
	}
	log.Debugf("loaded spec: %+v", Spec)
	return nil
}

func (p *{{ .Name }}Plugin) Stats() ([]byte, error) {
	return nil, nil
}
`
