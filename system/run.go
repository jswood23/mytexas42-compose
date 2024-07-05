package system

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

func Run(command string, args ...string) error {

	name := command

	if runtime.GOOS == "windows" {
		name = "cmd"
		args = append([]string{"/C"}, command+" "+strings.Join(args, " "))
	}

	println("Running command: " + name + " " + strings.Join(args, " "))

	cmd := exec.Command(name, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}

	return nil
}
