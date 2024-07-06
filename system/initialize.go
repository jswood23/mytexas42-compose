package system

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
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

// for testing
func tryPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		println("Path does not exist: " + path)
	} else {
		println("Path exists: " + path)
	}
}

func GetSSHKeyPath() string {
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
