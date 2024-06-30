package deployment

import (
	"fmt"
	"os/exec"
)

func deployAll() error {
	println("Deploying full stack.")
	_, err := exec.Command("cmd", "/C", "docker compose up -d").Output()

	if err != nil {
		return err
	}

	println("Deployment complete.")

	return nil
}

func deployEnvironment(environment string) error {
	println(fmt.Sprintf("Deploying to %s.", environment))
	_, err := exec.Command("cmd", "/C", fmt.Sprintf("docker compose build backend-%s", environment)).Output()

	if err != nil {
		return err
	}

	_, err = exec.Command("cmd", "/C", fmt.Sprintf("docker compose up --no-deps -d backend-%s", environment)).Output()

	if err != nil {
		return err
	}

	_, err = exec.Command("cmd", "/C", fmt.Sprintf("docker compose build frontend-%s", environment)).Output()

	if err != nil {
		return err
	}

	_, err = exec.Command("cmd", "/C", fmt.Sprintf("docker compose up --no-deps -d frontend-%s", environment)).Output()

	if err != nil {
		return err
	}

	return nil
}
