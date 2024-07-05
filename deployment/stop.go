package deployment

import (
	"mytexas42-compose/system"
)

func stopAll() error {
	println("Stopping full stack.")
	err := system.Run("docker", "compose down")

	if err != nil {
		return err
	}

	println("Stop complete.")

	return nil
}
