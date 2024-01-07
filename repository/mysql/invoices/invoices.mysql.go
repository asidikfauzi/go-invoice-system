package invoices

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"gorm.io/gorm"
)

type InvoicesMysql interface {
	Begin() *gorm.DB
	Commit() error
	Rollback() error
	GetAll(limit, offset int, orderBy string, request model.RequestInvoices) ([]model.GetInvoices, int64, error)
	FindById(invoiceId string) (model.GetInvoice, error)
	CheckExistsInvoiceId(invoiceId string) (bool, error)
	Create(invoice *domain.Invoices, items []domain.InvoiceHasItems) error
}
