package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"mytexas42-compose/deployment"
	"net/http"
	"os"
)

func main() {
	err := deployment.InitializeGit()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	err = r.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	r.GET("/health", getAppHealth)
	r.GET("/deploy/staging", deployment.DeployStaging)
	r.GET("/deploy/production", deployment.DeployProduction)
	r.GET("/deploy/all", deployment.DeployAll)
	r.GET("/stop/all", deployment.StopAll)
	r.GET("/stop/compose", deployment.StopCompose)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("CICD_PORT")

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func getAppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
