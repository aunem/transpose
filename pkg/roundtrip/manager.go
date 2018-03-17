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
const Dir = "./plugins/roundtrip"

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
func NewManager(c *config.Transpose) (*Manager, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	log.Debugf("files: %+v", files)
	m := Manager{
		Roundtrip: Roundtrip{},
	}
	rtp, err := loadRoundtrip(files, c.Roundtrip)
	if err != nil {
		return nil, err
	}
	m.Roundtrip = Roundtrip{Plugin: rtp, Config: c.Roundtrip.Spec}
	log.Debugf("manager: %+v", m)
	return &m, nil
}

func loadRoundtrip(files []os.FileInfo, plugin config.RoundtripPlugin) (rtp Plugin, err error) {
	path := soPath(files, plugin)
	if path == "" {
		path, err = resolve.ResolvePlugin(plugin.Name, plugin.Package)
		if err != nil {
			return rtp, err
		}
		path, err = resolve.BuildPlugin(plugin.Name, plugin.Package, resolve.RoundtripType)
		if err != nil {
			return rtp, err
		}
	}
	p, err := plug.Open(path)
	if err != nil {
		return rtp, err
	}
	symPlugin, err := p.Lookup("RoundtripPlugin")
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
	respCtx, err := m.Roundtrip.Plugin.Roundtrip(ctx, m.Roundtrip.Config)
	if err != nil {
		return nil, err
	}
	return respCtx, nil
}
