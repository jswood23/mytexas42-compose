package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mytexas42-compose/deployment"
	"mytexas42-compose/system"
	"net/http"
)

func main() {
	err := system.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	err = r.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	deployment.DeployAllNoContext()

	r.GET("/health", getAppHealth)
	r.GET("/deploy/staging", deployment.DeployStaging)
	r.GET("/deploy/production", deployment.DeployProduction)
	r.GET("/deploy/all", deployment.DeployAll)
	r.GET("/stop/all", deployment.StopAll)
	r.GET("/stop/compose", deployment.StopCompose)

	err = r.Run(":" + system.GetPort())
	if err != nil {
		log.Fatal(err)
	}
}

func getAppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
