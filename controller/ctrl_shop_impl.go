package controller

import (
	"net/http"
	"personal-finance/models/web"
	"personal-finance/services"
	"personal-finance/utils"
	"strconv"

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
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	utils.PanicIfError(err)

	ShopResponse := controller.ShopService.GetById(c, idInt)
	WebResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ShopResponse,
	}
	c.JSON(http.StatusOK, WebResponse)
}

func (controller *ShopControllerImpl) GetAll(c *gin.Context) {

}

func (controller *ShopControllerImpl) Update(c *gin.Context) {

}

func (controller *ShopControllerImpl) Delete(c *gin.Context) {

}
