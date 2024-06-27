package deployment

import (
	"github.com/gin-gonic/gin"
)

func DeployStaging(c *gin.Context) {
	deploy(c, "staging")
}

func DeployProduction(c *gin.Context) {
	deploy(c, "production")
}
