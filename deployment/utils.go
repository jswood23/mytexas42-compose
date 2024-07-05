package deployment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func updateAndDeployEnvironment(environment string) error {
	err := updateGit(environment)

	if err != nil {
		return err
	}

	err = deployEnvironment(environment)

	if err != nil {
		return err
	}

	return nil
}

func writeError(c *gin.Context, err error) {
	println(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "error",
		"reason": err.Error(),
	})
}

func checkPermissions(c *gin.Context) error {
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	userPassword := c.Query("password")

	if userPassword != adminPassword {
		return errors.New("Invalid password.")
	}

	return nil
}
