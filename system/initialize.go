package system

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

var sshPassphrase string
var port string
var sshKeyPath string

func Initialize() error {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file: " + err.Error())
	}
	port = os.Getenv("CICD_PORT")
	sshKeyPath = os.Getenv("SSH_KEY_PATH")

	// throw if environment variables are missing
	if port == "" || sshKeyPath == "" {
		return errors.New("missing environment variables")
	}

	// check OS args
	if len(os.Args) < 2 {
		return errors.New("missing password argument")
	}
	sshPassphrase = os.Args[1]

	return nil
}

func GetSSHPassphrase() string {
	return sshPassphrase
}

func GetPort() string {
	return port
}

func GetSSHKeyPath() string {
	if _, err := os.Stat(sshKeyPath); os.IsNotExist(err) {
		fmt.Printf("SSH key path %s does not exist.\n", sshKeyPath)
	}

	// Replace ~ with the home directory path
	if strings.HasPrefix(sshKeyPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			//return errors.New("error retrieving home directory: " + err.Error())
			println(err)
		}
		sshKeyPath = filepath.Join(homeDir, sshKeyPath[1:])
	}

	if _, err := os.Stat(sshKeyPath); os.IsNotExist(err) {
		fmt.Printf("SSH key path %s does not exist.\n", sshKeyPath)
	}

	return sshKeyPath
}

func GetCodePaths(environment string) (string, string) {
	var backendPath, frontendPath string
	if environment == "staging" {
		backendPath = os.Getenv("BACKEND_STAGING_REPO")
		frontendPath = os.Getenv("FRONTEND_STAGING_REPO")
	} else {
		backendPath = os.Getenv("BACKEND_PRODUCTION_REPO")
		frontendPath = os.Getenv("FRONTEND_PRODUCTION_REPO")
	}

	return backendPath, frontendPath
}
