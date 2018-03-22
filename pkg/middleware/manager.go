package middleware

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

// Dir is the directory that holds the .so middleware files
const Dir = "./plugins/middleware"

// Manager manages middleware
type Manager struct {
	RequestMiddlewares  []Middleware
	ResponseMiddlewares []Middleware
}

// Middleware holds the plugin and variables to be executed
type Middleware struct {
	Plugin Plugin
	Config interface{}
}

// NewManager returns a new middleware manager
func NewManager(c *config.Transpose) (*Manager, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	log.Debugf("files: %+v", files)
	m := Manager{
		RequestMiddlewares:  []Middleware{},
		ResponseMiddlewares: []Middleware{},
	}
	for _, plugin := range c.Spec.Middleware.Request {
		mw, err := loadMiddleware(files, plugin)
		if err != nil {
			return nil, err
		}
		m.RequestMiddlewares = append(m.RequestMiddlewares, mw)
	}
	for _, plugin := range c.Spec.Middleware.Response {
		mw, err := loadMiddleware(files, plugin)
		if err != nil {
			return nil, err
		}
		m.ResponseMiddlewares = append(m.ResponseMiddlewares, mw)
	}
	log.Debugf("manager: %+v", m)
	return &m, nil
}

func loadMiddleware(files []os.FileInfo, plugin config.MiddlewarePlugin) (mw Middleware, err error) {
	path := soPath(files, plugin)
	if path == "" {
		path, err = resolve.ResolvePlugin(plugin.Name, plugin.Package)
		if err != nil {
			return mw, err
		}
		path, err = resolve.BuildPlugin(plugin.Name, path, resolve.MiddlewareType)
		if err != nil {
			return mw, err
		}
	}
	p, err := plug.Open(path)
	if err != nil {
		return mw, err
	}
	symPlugin, err := p.Lookup("MiddlewarePlugin")
	if err != nil {
		return mw, err
	}

	mwp, ok := symPlugin.(Plugin)
	if !ok {
		return mw, fmt.Errorf("could not cast plugin type")
	}

	// symSpec, err := p.Lookup("Spec")  //TODO: see if this can be dynamically set
	// if err != nil {
	// 	return mw, err
	// }

	return Middleware{Plugin: mwp, Config: plugin.Spec}, nil
}

func soPath(fl []os.FileInfo, plugin config.MiddlewarePlugin) (path string) {
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

// ExecRequestStack executes all request middleware plugins
func (m *Manager) ExecRequestStack(ctx context.Request) (context.Request, error) {
	var err error
	for _, mw := range m.RequestMiddlewares {
		log.Debugf("executing middleware: %+v", mw)
		ctx, err = mw.Plugin.ProcessRequest(ctx, mw.Config)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

// ExecResponseStack executes all response middleware plugins
func (m *Manager) ExecResponseStack(ctx context.Response) (context.Response, error) {
	var err error
	for _, mw := range m.ResponseMiddlewares {
		log.Debugf("executing middleware: %+v", mw)
		ctx, err = mw.Plugin.ProcessResponse(ctx, mw.Config)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}
