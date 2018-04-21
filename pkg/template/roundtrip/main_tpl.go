package roundtrip

var mainTpl = `package main

import (
	"fmt"
	"path"

	"github.com/aunem/transpose/pkg/context"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type {{ .Name }}Plugin struct {}

// RoundtripPlugin exports the plugin struct
var RoundtripPlugin {{ .Name }}Plugin

// Spec exports the spec data
var Spec {{ .Name }}Spec

func main() {}

func (s *{{ .Name }}Plugin) Roundtrip(req context.Request) (context.Response, error) {
	// your roundtrip code goes here
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
