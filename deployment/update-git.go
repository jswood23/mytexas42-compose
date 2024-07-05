package deployment

import (
	"fmt"
	"mytexas42-compose/system"
	"os"
)

func updateGit(environment string) error {
	println(fmt.Sprintf("Updating %s from latest version.", environment))

	var backendPath string
	var frontendPath string
	if environment == "staging" {
		backendPath = os.Getenv("BACKEND_STAGING_REPO")
		frontendPath = os.Getenv("FRONTEND_STAGING_REPO")
	} else {
		backendPath = os.Getenv("BACKEND_PRODUCTION_REPO")
		frontendPath = os.Getenv("FRONTEND_PRODUCTION_REPO")
	}

	err := system.Run("git", fmt.Sprintf("-C %s/ pull", backendPath))

	if err != nil {
		return err
	}

	err = system.Run("git", fmt.Sprintf("-C %s/ pull", frontendPath))

	if err != nil {
		return err
	}

	return nil
}
