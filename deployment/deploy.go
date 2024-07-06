package deployment

import (
	"fmt"
	"mytexas42-compose/system"
)

func deployAll() error {
	println("Deploying full stack.")
	err := system.Run("docker", "compose", "up", "-d")

	if err != nil {
		return err
	}

	println("Deployment complete.")

	return nil
}

func deployEnvironment(environment string) error {
	println(fmt.Sprintf("Deploying to %s.", environment))
	err := system.Run("docker", fmt.Sprintf("compose", "build", "backend-%s", environment))

	if err != nil {
		return err
	}

	err = system.Run("docker", fmt.Sprintf("compose", "up", "--no-deps", "-d", "backend-%s", environment))

	if err != nil {
		return err
	}

	err = system.Run("docker", fmt.Sprintf("compose", "build", "frontend-%s", environment))

	if err != nil {
		return err
	}

	err = system.Run("docker", fmt.Sprintf("compose up --no-deps -d frontend-%s", environment))

	if err != nil {
		return err
	}

	println("Deployment complete.")

	return nil
}
