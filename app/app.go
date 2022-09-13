package app

import (
	"personal-finance/routes"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
