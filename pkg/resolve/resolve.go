package plugin

import (
	"bufio"
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

// ExecCommand executes the dep command to resolve the plugin
func ExecCommand(pkg string) {
	// docker build current directory
	cmdName := "dep"
	cmdArgs := []string{"ensure", "-v", "-add", pkg}

	cmd := exec.Command(cmdName, cmdArgs...)
	log.Debugf("cmd: %+v", cmd)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating StdoutPipe for Cmd: ", err)
	}

	done := make(chan struct{})
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			log.Infof("resolve out | %s\n", scanner.Text())
		}
		done <- struct{}{}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal("Error starting Cmd: ", err)
	}

	<-done
	err = cmd.Wait()
	if err != nil {
		log.Fatal("Error waiting for Cmd: ", err)
	}
}

// ResolvePlugin pulls a remote plugin local
func ResolvePlugin(name, pkg string) (path string, err error) {
	log.Debug("resolving plugin: ", pkg)
	ExecCommand(pkg)
	pkgPath := filepath.Join("./vendor", pkg)
	log.Debugf("package path: %+v", pkgPath)
	return pkgPath, nil
}

// BuildPlugin pulls a remote plugin local
func BuildPlugin(name, pkgPath string, typ Type) (path string, err error) {
	log.Debug("building plugin...")
	soPath := filepath.Join(Dir, string(typ), name, fmt.Sprintf("%s.so", name))
	cmd := exec.Command("go", "build", "-buildmode=plugin", fmt.Sprintf("-o %s", soPath), pkgPath)
	log.Debugf("cmd: %+v", cmd)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not build plugin %s: %s", pkgPath, err)
	}
	return soPath, nil
}
