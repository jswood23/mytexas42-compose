package deployment

import (
	"mytexas42-compose/system"
	"runtime"
)

func InitializeGit() error {
	println("Initializing git configuration.")

	var err error

	if runtime.GOOS != "windows" {
		// start ssh agent
		err = system.Run("eval", "$(ssh-agent -s)")
		if err != nil {
			return err
		}
	}

	// Set git configuration for dubious ownership error
	err = system.Run("git", "config", "--global", "core.safecrlf", "false")
	if err != nil {
		return err
	}

	// Set git configuration for dubious ownership error
	err = system.Run("git", "config", "--global", "core.autocrlf", "false")
	if err != nil {
		return err
	}

	return nil
}
