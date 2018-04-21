package template

import (
	"fmt"
	"os"
	tpl "text/template"

	list "github.com/aunem/transpose/pkg/template/listener"
	mw "github.com/aunem/transpose/pkg/template/middleware"
	rt "github.com/aunem/transpose/pkg/template/roundtrip"
)

//TODO: test this

// TemplateDir represents a templated directory
type TemplateDir string

var middlewaredir TemplateDir = "pkg/template/middleware"
var roundtripdir TemplateDir = "pkg/template/roundtrip"
var listenerdir TemplateDir = "pkg/template/listener"

// Plugin represents the template data for a plugin
type Plugin struct {
	Name, Pkg, Typ string
}

// NewPlugin creates a new plugin
func NewPlugin(name, pkg, typ string) *Plugin {
	return &Plugin{
		Name: name,
		Pkg:  pkg,
		Typ:  typ,
	}
}

// Template will take the plugin type and template the appropriate files
func (p *Plugin) Template() error {
	switch p.Typ {
	case "middleware":
		for k, v := range mw.Manifest {
			t, err := tpl.New("tpl").Parse(v)
			if err != nil {
				return err
			}
			f, err := os.Create(k)
			if err != nil {
				return err
			}
			defer f.Close()
			t.Execute(f, p)
			if err != nil {
				return err
			}
		}
	case "listener":
		for k, v := range list.Manifest {
			t, err := tpl.New("tpl").Parse(v)
			if err != nil {
				return err
			}
			f, err := os.Create(k)
			if err != nil {
				return err
			}
			defer f.Close()
			t.Execute(f, p)
			if err != nil {
				return err
			}
		}
	case "roundtrip":
		for k, v := range rt.Manifest {
			t, err := tpl.New("tpl").Parse(v)
			if err != nil {
				return err
			}
			f, err := os.Create(k)
			if err != nil {
				return err
			}
			defer f.Close()
			t.Execute(f, p)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("plugin typ uknown: %+v", p.Typ)
	}
	return nil
}
