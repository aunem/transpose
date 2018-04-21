package listener

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	plug "plugin"

	"github.com/aunem/transpose/config"
	"github.com/aunem/transpose/pkg/middleware"
	resolve "github.com/aunem/transpose/pkg/resolve"
	"github.com/aunem/transpose/pkg/roundtrip"
	log "github.com/sirupsen/logrus"
)

// Bin is the directory that holds the .so Listener files
const Bin = "./bin/listener"

// Manager manages Listener
type Manager struct {
	Listener   Listener
	Middleware *middleware.Manager
	Roundtrip  *roundtrip.Manager
	Spec       config.TransposeSpec
}

// Listener holds the plugin and variables to be executed
type Listener struct {
	Plugin Plugin
	Vars   map[string]string
	Spec   interface{}
}

// NewManager returns a new Listener manager
func NewManager(spec config.TransposeSpec, mw *middleware.Manager, rt *roundtrip.Manager) (*Manager, error) {
	files, err := ioutil.ReadDir(Bin)
	if err != nil {
		return nil, err
	}
	log.Debugf("files: %+v", files)
	m := Manager{
		Listener: Listener{},
	}
	log.Debug("loading plugin...")
	rtp, err := loadListener(files, spec)
	if err != nil {
		return nil, err
	}
	log.Debug("running init function...")
	err = rtp.Init(spec.Listener.Spec)
	if err != nil {
		return nil, err
	}
	m.Listener = Listener{Plugin: rtp, Spec: spec}
	m.Middleware = mw
	m.Roundtrip = rt
	m.Spec = spec
	log.Debugf("manager: %+v", m)
	return &m, nil
}

func loadListener(files []os.FileInfo, spec config.TransposeSpec) (rtp Plugin, err error) {
	plugin := spec.Listener
	path := soPath(files, plugin)
	log.Debug(".so file: ", path)
	if path == "" {
		log.Debug(".so file not found, resolving plugin...")
		if spec.LocalBuild {
			log.Info("resolving locally")
			path, err = resolve.ResolveLocal(plugin.Name, plugin.Package)
			if err != nil {
				log.Info("could not resolve plugin locally, trying remote...")
				path, err = resolve.ResolveRemote(plugin.Name, plugin.Package)
				if err != nil {
					return rtp, err
				}
			}
		} else {
			log.Info("resolving remote")
			path, err = resolve.ResolveRemote(plugin.Name, plugin.Package)
			if err != nil {
				return rtp, err
			}
		}
		path, err = resolve.BuildPlugin(plugin.Name, path, resolve.ListenerType)
		if err != nil {
			return rtp, err
		}
	}
	p, err := plug.Open(path)
	if err != nil {
		return rtp, err
	}
	symPlugin, err := p.Lookup("ListenerPlugin")
	if err != nil {
		return rtp, err
	}
	log.Debugf("sym plugin: %+v", symPlugin)
	var ok bool
	rtp, ok = symPlugin.(Plugin)
	if !ok {
		return rtp, fmt.Errorf("could not cast plugin type")
	}
	return rtp, nil
}

func soPath(fl []os.FileInfo, plugin config.ListenerPlugin) (path string) {
	log.Debug("checking for .so file...")
	for _, fileInfo := range fl {
		if fileInfo.IsDir() {
			if plugin.Name == fileInfo.Name() {
				fp := filepath.Join(Bin, fileInfo.Name(), fmt.Sprintf("%s.so", fileInfo.Name()))
				if _, err := os.Stat(fp); err == nil {
					return fp
				}
			}
		}
	}
	return ""
}

// ExecListener executes the Listener plugin
func (m *Manager) ExecListener() error {
	log.Debugf("executing Listener: %+v", m.Listener)
	err := m.Listener.Plugin.Listen(m.Middleware, m.Roundtrip)
	return err
}
