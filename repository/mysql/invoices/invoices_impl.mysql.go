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

func (m *Invoices) Begin() *gorm.DB {
	return m.DB.Begin()
}

func (m *Invoices) Commit() error {
	return m.DB.Commit().Error
}

func (m *Invoices) Rollback() error {
	return m.DB.Rollback().Error
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
		query = query.Order("invoices.created_at " + orderBy)
	} else {
		query = query.Order("invoices.created_at ASC")
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

func (m *Invoices) FindById(invoiceId string) (model.GetInvoice, error) {
	var (
		invoice         domain.Invoices
		invoiceHasItems []domain.InvoiceHasItems
		items           []model.GetItem
		data            model.GetInvoice
	)

	if err := m.DB.Preload("Customers").
		Joins("JOIN customers c ON invoices.customer_id = c.id_customer").
		Where("invoices.id_invoice = ?", invoiceId).
		Where("invoices.deleted_at IS NULL").
		First(&invoice).Error; err != nil {
		return data, err
	}

	if err := m.DB.Preload("Items").
		Preload("Items.Types").
		Joins("JOIN items i ON invoice_has_items.item_id = i.id_item").
		Joins("JOIN types t ON i.type_id = t.id_type").
		Where("invoice_has_items.invoice_id = ?", invoiceId).
		Where("invoice_has_items.deleted_at IS NULL").
		Find(&invoiceHasItems).Error; err != nil {
		return data, err
	}

	for _, ihi := range invoiceHasItems {
		item := model.GetItem{
			IDItem:       ihi.Items.IDItem,
			ItemName:     ihi.Items.ItemName,
			ItemQuantity: ihi.Items.ItemQuantity,
			ItemPrice:    ihi.Items.ItemPrice,
			TypeID:       ihi.Items.Types.IDType,
			TypeName:     ihi.Items.Types.TypeName,
		}

		items = append(items, item)
	}

	data = model.GetInvoice{
		IDInvoice:         invoice.IDInvoice,
		InvoiceID:         invoice.InvoiceID,
		InvoiceSubject:    invoice.InvoiceSubject,
		InvoiceIssueDate:  invoice.InvoiceIssueDate,
		InvoiceDueDate:    invoice.InvoiceDueDate,
		InvoiceTotalItem:  invoice.InvoiceTotalItem,
		InvoiceSubTotal:   invoice.InvoiceSubTotal,
		InvoiceTax:        invoice.InvoiceTax,
		InvoiceGrandTotal: invoice.InvoiceGrandTotal,
		InvoiceStatus:     invoice.InvoiceStatus,
		Customer: model.GetCustomer{
			IDCustomer:      invoice.Customers.IDCustomer,
			CustomerName:    invoice.Customers.CustomerName,
			CustomerAddress: invoice.Customers.CustomerAddress,
		},
		Items: items,
	}

	return data, nil
}

func (m *Invoices) CheckExistsInvoiceId(invoiceId string) (bool, error) {
	var (
		invoice domain.Invoices
		err     error
	)

	if err = m.DB.Where("invoice_id = ?", invoiceId).
		Where("deleted_at IS NULL").
		First(&invoice).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (m *Invoices) FindInvoiceHasItems(invoiceId string) ([]model.GetInvoiceHasItem, error) {
	var (
		invoiceHasItem []domain.InvoiceHasItems
		data           []model.GetInvoiceHasItem
		err            error
	)

	if err = m.DB.Where("invoice_id = ?", invoiceId).
		Where("deleted_at IS NULL").
		Find(&invoiceHasItem).Error; err != nil {
		return data, err
	}

	for _, ihi := range invoiceHasItem {
		inHasIt := model.GetInvoiceHasItem{
			InvoiceID: ihi.InvoiceID,
			ItemID:    ihi.ItemID,
			Quantity:  ihi.Quantity,
		}

		data = append(data, inHasIt)
	}

	return data, nil
}
