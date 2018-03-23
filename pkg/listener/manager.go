package listener

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	plug "plugin"

	"github.com/aunem/transpose/config"
	resolve "github.com/aunem/transpose/pkg/resolve"
	log "github.com/sirupsen/logrus"
)

// Dir is the directory that holds the .so Listener files
const Dir = "./bin/listener"

// Manager manages Listener
type Manager struct {
	Listener Listener
	Spec     config.TransposeSpec
}

// Listener holds the plugin and variables to be executed
type Listener struct {
	Plugin Plugin
	Vars   map[string]string
	Spec   interface{}
}

// NewManager returns a new Listener manager
func NewManager(spec config.TransposeSpec) (*Manager, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	log.Debugf("files: %+v", files)
	m := Manager{
		Listener: Listener{},
	}
	rtp, err := loadListener(files, spec.Listener)
	if err != nil {
		return nil, err
	}
	m.Listener = Listener{Plugin: rtp, Spec: spec}
	m.Spec = spec
	log.Debugf("manager: %+v", m)
	return &m, nil
}

func loadListener(files []os.FileInfo, plugin config.ListenerPlugin) (rtp Plugin, err error) {
	path := soPath(files, plugin)
	log.Debug(".so file: ", path)
	if path == "" {
		log.Debug(".so file not found, resolving plugin...")
		path, err = resolve.ResolvePlugin(plugin.Name, plugin.Package)
		if err != nil {
			return rtp, err
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
				fp := filepath.Join(Dir, fileInfo.Name(), fmt.Sprintf("%s.so", fileInfo.Name()))
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
	err := m.Listener.Plugin.Listen(m.Spec)
	return err
}
