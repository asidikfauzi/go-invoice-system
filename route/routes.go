package route

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/controller/types"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	TypeService types.TypeController `inject:"controller_type_master"`
}

func InitPackage() *RouteService {
	return &RouteService{
		TypeService: &types.MasterTypes{},
	}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	endpoint := router.Group("/inv")
	{
		types := endpoint.Group("/type")
		{
			types.GET("", r.TypeService.GetAllTypes)
			types.GET("/:typeId", r.TypeService.FindTypeById)
			types.POST("", r.TypeService.CreateType)
			types.PATCH("/:typeId", r.TypeService.UpdateType)
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
