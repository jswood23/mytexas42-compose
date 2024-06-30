package deployment

import (
	"fmt"
	"os"
	"os/exec"
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

	_, err := exec.Command("cmd", "/C", fmt.Sprintf("git -C %s/ pull", backendPath)).Output()

	if err != nil {
		return err
	}

	_, err = exec.Command("cmd", "/C", fmt.Sprintf("git -C %s/ pull", frontendPath)).Output()

	if err != nil {
		return err
	}

	return nil
}
