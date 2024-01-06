package types

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"time"
)

type TypesService interface {
	GetAllTypes(c *gin.Context, pageParam, limitParam, orderByParam, typeName string, startTime time.Time) ([]model.Types, helper.Paginate, error)
	FindTypeById(c *gin.Context, typeId string, startTime time.Time) (model.Types, error)
	CreateType(c *gin.Context, request model.RequestType, startTime time.Time) (string, error)
}
