package invoices

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Invoices struct {
	DB *gorm.DB
}

func NewInvoiceMysql(conn *gorm.DB) InvoicesMysql {
	return &Invoices{
		DB: conn,
	}
}

func (m *Invoices) GetAll(limit, offset int, orderBy string, request model.RequestInvoices) ([]model.GetInvoices, int64, error) {
	var (
		invoices   []domain.Invoices
		data       []model.GetInvoices
		totalCount int64
	)

	query := m.DB.Where("invoices.deleted_at IS NULL")

	if request.InvoiceID != "" {
		query = query.Where("invoices.invoice_id LIKE ?", "%"+request.InvoiceID+"%")
	}

	if request.InvoiceIssueDate != "" {
		issueDate, err := time.Parse("2006-01-02", request.InvoiceIssueDate)
		if err != nil {
			return nil, totalCount, err
		}
		query = query.Where("DATE(invoices.invoice_issue_date) = ?", issueDate.Format("2006-01-02"))
	}

	if request.InvoiceDueDate != "" {
		dueDate, err := time.Parse("2006-01-02", request.InvoiceDueDate)
		if err != nil {
			return nil, totalCount, err
		}
		query = query.Where("DATE(invoices.invoice_due_date) = ?", dueDate.Format("2006-01-02"))
	}

	if request.InvoiceSubject != "" {
		query = query.Where("invoices.invoice_subject LIKE ?", "%"+request.InvoiceSubject+"%")
	}

	if request.InvoiceTotalItem != 0 {
		totalItem := strconv.Itoa(request.InvoiceTotalItem)
		query = query.Where("invoices.invoice_total_item LIKE ?", "%"+totalItem+"%")
	}

	if request.CustomerName != "" {
		query = query.Where("c.customer_name LIKE ?", "%"+request.CustomerName+"%")
	}

	if request.InvoiceStatus != "" {
		query = query.Where("invoices.invoice_status = ?", request.InvoiceStatus)
	}

	if orderBy != "" {
		query = query.Order("invoices.invoice_id " + orderBy)
	} else {
		query = query.Order("invoices.invoice_id ASC")
	}

	if err := query.Preload("Customers").
		Joins("JOIN customers c ON invoices.customer_id = c.id_customer").
		Offset(offset).
		Limit(limit).
		Find(&invoices).Error; err != nil {
		return nil, totalCount, err
	}

	if err := query.Model(&domain.Invoices{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	for _, i := range invoices {
		invoice := model.GetInvoices{
			IDInvoice:        i.IDInvoice,
			InvoiceID:        i.InvoiceID,
			InvoiceSubject:   i.InvoiceSubject,
			InvoiceIssueDate: i.InvoiceIssueDate,
			InvoiceDueDate:   i.InvoiceDueDate,
			InvoiceTotalItem: i.InvoiceTotalItem,
			InvoiceStatus:    i.InvoiceStatus,
			CustomerName:     i.Customers.CustomerName,
		}
		data = append(data, invoice)
	}

	return data, totalCount, nil
}
