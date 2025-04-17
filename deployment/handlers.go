package deployment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DeployStaging(c *gin.Context) {
	err := updateAndDeployEnvironment("staging")

	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}

func DeployProduction(c *gin.Context) {
	err := updateAndDeployEnvironment("production")

	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}

func DeployAll(c *gin.Context) {
	err := checkPermissions(c)

	if err != nil {
		writeError(c, err)
		return
	}

	err = updateGit("staging")

	if err != nil {
		writeError(c, err)
		return
	}

	err = updateGit("production")

	if err != nil {
		writeError(c, err)
		return
	}

	err = deployAll()

	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}

func DeployAllNoContext() {
	_ = updateGit("staging")
	_ = updateGit("production")
	_ = deployAll()
}

func StopAll(c *gin.Context) {
	err := checkPermissions(c)

	if err != nil {
		writeError(c, err)
		return
	}

	err = stopAll()

	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}

func StopCompose(c *gin.Context) {
	err := checkPermissions(c)

	if err != nil {
		writeError(c, err)
		return
	}

	// stop server here
	defer os.Exit(0)

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}
