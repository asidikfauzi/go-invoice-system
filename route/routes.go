package route

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
}

func InitPackage() *RouteService {
	return &RouteService{}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
