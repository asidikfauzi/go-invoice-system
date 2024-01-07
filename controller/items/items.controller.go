package items

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/service/items"
)

type ItemController interface {
	GetAllItems(c *gin.Context)
	FindItemById(c *gin.Context)
	CreateItem(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}

type MasterItems struct {
	ItemService items.ItemsService `inject:"items_service"`
}
