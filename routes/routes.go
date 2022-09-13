package routes

import (
	"personal-finance/services"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", services.Dashboard)
}
