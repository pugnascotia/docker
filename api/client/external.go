package client

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// CmdExternalCommand runs an external command from $PATH, which must be
// prefixed with "docker-". All command line options are passed to the
// command
//
// Usage: docker COMMAND [OPTIONS]
func (cli *DockerCli) CmdExternalCommand(args ...string) error {
	path, _ := exec.LookPath("docker-" + args[0])

	// By convention, args[0] should contain the filename associated with
	// the file being executed.
	var cmdArgs []string
	if len(args) == 0 {
		cmdArgs = []string{path}
	} else {
		cmdArgs = append([]string{path}, args[1:]...)
	}

	if err := syscall.Exec(path, cmdArgs, os.Environ()); err != nil {
		return fmt.Errorf("Unable to execute docker-%s - %s", args[0], err)
	}
	return nil
}
