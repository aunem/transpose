package middleware

var mainTpl = `package main

import (
	"fmt"
	"path"

	"github.com/aunem/transpose/pkg/context"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type {{ .Name }}Plugin struct {}

// MiddlewarePlugin exports the plugin struct
var MiddlewarePlugin {{ .Name }}Plugin

// Spec exports the spec data
var Spec {{ .Name }}Spec

func main() {}

func (p *{{ .Name }}Plugin) ProcessRequest(req context.Request) (context.Request, error) {
    return nil, nil
}

func (p *{{ .Name }}Plugin) ProcessResponse(resp context.Response) (context.Response, error) {
	return nil, nil
}

func (p *{{ .Name }}Plugin) LoadSpec(spec interface{}) error {
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

func (p *{{ .Name }}Plugin) Init() error {
	return nil
}

func (p *{{ .Name }}Plugin) Stats() ([]byte, error) {
	return nil, nil
}
`
