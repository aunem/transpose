package roundtrip

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	plug "plugin"

	"github.com/aunem/transpose/config"
	"github.com/aunem/transpose/pkg/context"
	resolve "github.com/aunem/transpose/pkg/resolve"
	log "github.com/sirupsen/logrus"
)

// Dir is the directory that holds the .so Roundtrip files
const Dir = "./bin/roundtrip"

// Manager manages roundtrip
type Manager struct {
	Roundtrip Roundtrip
}

// Roundtrip holds the plugin and variables to be executed
type Roundtrip struct {
	Plugin Plugin
	Config interface{}
}

// NewManager returns a new Roundtrip manager
func NewManager(spec config.TransposeSpec) (*Manager, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	log.Debugf("files: %+v", files)
	m := Manager{
		Roundtrip: Roundtrip{},
	}
	log.Debug("loading plugin...")
	rtp, err := loadRoundtrip(files, spec)
	if err != nil {
		return nil, err
	}
	log.Debug("loading spec...")
	err = rtp.LoadSpec(spec.Roundtrip.Spec)
	if err != nil {
		return nil, err
	}
	log.Debug("running init...")
	err = rtp.Init()
	if err != nil {
		return nil, err
	}
	m.Roundtrip = Roundtrip{Plugin: rtp, Config: spec.Roundtrip.Spec}
	log.Debugf("manager: %+v", m)
	return &m, nil
}

func loadRoundtrip(files []os.FileInfo, spec config.TransposeSpec) (rtp Plugin, err error) {
	plugin := spec.Roundtrip
	path := soPath(files, plugin)
	if path == "" {
		if spec.LocalBuild {
			path, err = resolve.ResolveLocal(plugin.Name, plugin.Package)
			if err != nil {
				return rtp, err
			}
		} else {
			path, err = resolve.ResolveRemote(plugin.Name, plugin.Package)
			if err != nil {
				return rtp, err
			}
		}
		path, err = resolve.BuildPlugin(plugin.Name, path, resolve.RoundtripType)
		if err != nil {
			return rtp, err
		}
	}
	p, err := plug.Open(path)
	if err != nil {
		return rtp, err
	}
	log.Debug("loading plugin: ", path)
	symPlugin, err := p.Lookup("RoundtripPlugin")
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

func soPath(fl []os.FileInfo, plugin config.RoundtripPlugin) (path string) {
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

// ExecRoundtrip executes the roundtrip plugin
func (m *Manager) ExecRoundtrip(ctx context.Request) (context.Response, error) {
	var err error
	log.Debugf("executing roundtrip: %+v", m.Roundtrip)
	respCtx, err := m.Roundtrip.Plugin.Roundtrip(ctx)
	if err != nil {
		return nil, err
	}
	return respCtx, nil
}
