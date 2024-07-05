package system

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
)

func Run(command string) error {
	println("Running command: " + command)

	name := "sudo"
	args := []string{command}

	if runtime.GOOS == "windows" {
		name = "cmd"
		args = append([]string{"/C"}, args...)
	}

	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}

	return nil
}
