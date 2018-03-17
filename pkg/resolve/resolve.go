package plugin

import (
	"fmt"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Dir is the plugin directory
const Dir = "./plugins/"

// Type is an enum for a plugin
type Type string

const (
	// ListenerType is a listener plugin
	ListenerType Type = "listener"
	// MiddlewareType is a middleware plugin
	MiddlewareType Type = "middleware"
	// RoundtripType is a server plugin
	RoundtripType Type = "roundtrip"
)

// ResolvePlugin pulls a remote plugin local
func ResolvePlugin(name, pkg string) (path string, err error) {
	cmd := exec.Command("dep", "ensure", "-add", pkg)
	log.Debugf("cmd: %+v", cmd)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not fetch plugin %s: %s", pkg, err)
	}
	pkgPath := filepath.Join("./vendor", pkg)
	log.Debugf("package path: %+v", pkgPath)
	return pkgPath, nil
}

// BuildPlugin pulls a remote plugin local
func BuildPlugin(name, pkgPath string, typ Type) (path string, err error) {
	soPath := filepath.Join(Dir, string(typ), name, fmt.Sprintf("%s.so", name))
	cmd := exec.Command("go", "build", "-buildmode=plugin", fmt.Sprintf("-o %s", soPath), pkgPath)
	log.Debugf("cmd: %+v", cmd)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not build plugin %s: %s", pkgPath, err)
	}
	return soPath, nil
}
