package invoices

import (
	"go-invoice-system/model"
	"gorm.io/gorm"
)

type InvoicesMysql interface {
	Begin() *gorm.DB
	Commit() error
	Rollback() error
	GetAll(limit, offset int, orderBy string, request model.RequestInvoices) ([]model.GetInvoices, int64, error)
	FindById(invoiceId string) (model.GetInvoice, error)
	CheckExistsInvoiceId(invoiceId string) (bool, error)
	FindInvoiceHasItems(invoiceId string) ([]model.GetInvoiceHasItem, error)
}
