package types

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/service/types"
)

type TypeController interface {
	GetAllTypes(c *gin.Context)
	FindTypeById(c *gin.Context)
	CreateType(c *gin.Context)
	UpdateType(c *gin.Context)
	DeleteType(c *gin.Context)
}

type MasterTypes struct {
	TypeService types.TypesService `inject:"types_service"`
}
