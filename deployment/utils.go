package deployment

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func outputError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "error",
		"reason": err.Error(),
	})
}
