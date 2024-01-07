package invoices

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/service/invoices"
)

type InvoiceController interface {
	GetAllInvoices(c *gin.Context)
	FindInvoiceById(c *gin.Context)
	CreateInvoice(c *gin.Context)
	UpdateInvoice(c *gin.Context)
}

type MasterInvoices struct {
	InvoiceService invoices.InvoicesService `inject:"invoices_service"`
}
