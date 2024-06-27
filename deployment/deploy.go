package deployment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func deploy(c *gin.Context, environment string) {
	println(fmt.Sprintf("Deploying to %s.", environment))
	_, err := exec.Command("cmd", "/C", fmt.Sprintf("docker compose build backend-%s", environment)).Output()

	if err != nil {
		outputError(c, err)
		return
	}

	_, err = exec.Command("cmd", "/C", fmt.Sprintf("docker compose up --no-deps -d backend-%s", environment)).Output()

	if err != nil {
		outputError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "complete",
	})
}
