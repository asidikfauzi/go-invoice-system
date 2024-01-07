package invoices

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"time"
)

type InvoicesService interface {
	GetAllInvoices(c *gin.Context, pageParam, limitParam, orderByParam string, request model.RequestInvoices, startTime time.Time) ([]model.GetInvoices, helper.Paginate, error)
	FindInvoiceById(c *gin.Context, invoiceId string, startTime time.Time) (model.GetInvoice, error)
	CreateInvoice(c *gin.Context, request model.RequestInvoice, startTime time.Time) (string, error)
}
