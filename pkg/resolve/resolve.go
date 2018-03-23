package plugin

import (
	"bufio"
	"fmt"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Dir is the plugin directory
const Dir = "./bin/"

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

// AddPkg adds a package using dep
func AddPkg(pkg string) {
	cmdName := "dep"
	cmdArgs := []string{"ensure", "-v", "-add", pkg}
	ExecCommand(cmdName, cmdArgs)
}

// Update deps to the latest version
func Update() {
	cmdName := "dep"
	cmdArgs := []string{"ensure", "-v", "-update"}
	ExecCommand(cmdName, cmdArgs)
}

// ExecCommand executes the dep command to resolve the plugin
func ExecCommand(cmdName string, cmdArgs []string) {
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
	log.Info("resolving plugin: ", pkg)
	Update()
	AddPkg(pkg)
	pkgPath := filepath.Join("vendor", pkg)
	log.Infof("package path: %+v", pkgPath)
	return fmt.Sprintf("./%s", pkgPath), nil
}

// BuildPlugin pulls a remote plugin local
func BuildPlugin(name, pkgPath string, typ Type) (path string, err error) {
	log.Debug("building plugin...")
	soPath := filepath.Join(Dir, string(typ), name, fmt.Sprintf("%s.so", name))
	cmdName := "go"
	cmdArgs := []string{"build", "-buildmode=plugin", "-o", soPath, pkgPath}
	ExecCommand(cmdName, cmdArgs)
	return soPath, nil
}
