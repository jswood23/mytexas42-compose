package deployment

import (
	"fmt"
	"mytexas42-compose/system"
	"os"
)

func updateGit(environment string) error {
	println(fmt.Sprintf("Updating %s from latest version.", environment))

	var backendPath, frontendPath string
	if environment == "staging" {
		backendPath = os.Getenv("BACKEND_STAGING_REPO")
		frontendPath = os.Getenv("FRONTEND_STAGING_REPO")
	} else {
		backendPath = os.Getenv("BACKEND_PRODUCTION_REPO")
		frontendPath = os.Getenv("FRONTEND_PRODUCTION_REPO")
	}

	// Function to handle git pull and possible dubious ownership error
	handleGitPull := func(path string) error {
		return system.Run("git", "-C", path, "pull")
	}

	// Update backend repository
	if err := handleGitPull(backendPath); err != nil {
		return err
	}

	// Update frontend repository
	if err := handleGitPull(frontendPath); err != nil {
		return err
	}

	return nil
}
