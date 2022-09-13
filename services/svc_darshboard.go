package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":       "Dashboard",
		"description": "Welcome to the Dashboard",
	})
}
