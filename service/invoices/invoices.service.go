package invoices

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"time"
)

type InvoicesService interface {
	GetAllInvoices(c *gin.Context, pageParam, limitParam, orderByParam string, request model.RequestInvoices, startTime time.Time) ([]model.GetInvoices, helper.Paginate, error)
}
