package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutThisProject(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "A simple todo project using Gorm with Postgres, Gin with Logger middleware and Swagger Documentation"})
}
