package controller

import "github.com/gin-gonic/gin"

type ShopController interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
