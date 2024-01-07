package invoices

import (
	"go-invoice-system/model"
)

type InvoicesMysql interface {
	GetAll(limit, offset int, orderBy string, request model.RequestInvoices) ([]model.GetInvoices, int64, error)
}
