package items

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"time"
)

type ItemsService interface {
	GetAllItems(c *gin.Context, pageParam, limitParam, orderByParam, itemName string, startTime time.Time) ([]model.GetItem, helper.Paginate, error)
	FindItemById(c *gin.Context, itemId string, startTime time.Time) (model.GetItem, error)
	CreateItem(c *gin.Context, request model.RequestItem, startTime time.Time) (string, error)
	UpdateItem(c *gin.Context, request model.RequestItem, itemId string, startTime time.Time) (string, error)
	DeleteItem(c *gin.Context, itemId string, startTime time.Time) (string, error)
}
