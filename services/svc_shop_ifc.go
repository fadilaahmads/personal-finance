package services

import (
	"personal-finance/models/web"

	"github.com/gin-gonic/gin"
)

type ShopService interface {
	Create(c *gin.Context) web.ShopResponse
	GetById(c *gin.Context, id int) web.ShopResponse
	GetAll(c *gin.Context) web.ShopResponse
	Update(c *gin.Context, id int) web.ShopUpdateRequest
	Delete(c *gin.Context, id int)
}
