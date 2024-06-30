package deployment

import "os/exec"

func stopAll() error {
	println("Stopping full stack.")
	_, err := exec.Command("cmd", "/C", "docker compose down").Output()

	if err != nil {
		return err
	}

	println("Stop complete.")

	return nil
}
