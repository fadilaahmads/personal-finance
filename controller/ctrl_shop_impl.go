package controller

import (
	"personal-finance/services"

	"github.com/gin-gonic/gin"
)

type ShopControllerImpl struct {
	ShopService services.ShopService
}

func NewShopController(ShopService services.ShopService) ShopController {
	return &ShopControllerImpl{
		ShopService: ShopService,
	}
}

func (controller *ShopControllerImpl) Create(c *gin.Context) {

}

func (controller *ShopControllerImpl) GetById(c *gin.Context) {

}

func (controller *ShopControllerImpl) GetAll(c *gin.Context) {

}

func (controller *ShopControllerImpl) Update(c *gin.Context) {

}

func (controller *ShopControllerImpl) Delete(c *gin.Context) {

}
