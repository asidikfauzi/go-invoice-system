package invoices

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"go-invoice-system/repository/mysql/customers"
	"go-invoice-system/repository/mysql/invoices"
	"go-invoice-system/repository/mysql/items"
	"log"
	"math"
	"net/http"
	"time"
)

type Invoice struct {
	invoiceMysql  invoices.InvoicesMysql
	customerMysql customers.CustomersMysql
	itemMysql     items.ItemsMysql
}

func NewInvoiceService(im invoices.InvoicesMysql, cm customers.CustomersMysql, itm items.ItemsMysql) InvoicesService {
	return &Invoice{
		invoiceMysql:  im,
		customerMysql: cm,
		itemMysql:     itm,
	}
}

func (s *Invoice) GetAllInvoices(c *gin.Context, pageParam, limitParam, orderByParam string, request model.RequestInvoices, startTime time.Time) ([]model.GetInvoices, helper.Paginate, error) {
	var (
		dataInvoices []model.GetInvoices
		paginate     helper.Paginate
		totalData    int64
		err          error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return dataInvoices, paginate, err
	}

	dataInvoices, totalData, err = s.invoiceMysql.GetAll(limit, offset, orderByParam, request)
	if err != nil {
		log.Printf("error invoice service GetAll: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return dataInvoices, paginate, err
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	paginate = helper.Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return dataInvoices, paginate, nil
}

func (s *Invoice) FindInvoiceById(c *gin.Context, invoiceId string, startTime time.Time) (model.GetInvoice, error) {
	var (
		invoice model.GetInvoice
		err     error
	)

	invoice, err = s.invoiceMysql.FindById(invoiceId)
	if err != nil {
		err = fmt.Errorf("invoice_id '%s' not found", invoiceId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return invoice, err
	}

	return invoice, nil
}

func (s *Invoice) CreateInvoice(c *gin.Context, request model.RequestInvoice, startTime time.Time) (string, error) {
	var (
		invoiceHasItems []domain.InvoiceHasItems
		invoice         domain.Invoices
		getItem         model.GetItem
		updateItems     []domain.Items
		err             error
	)

	issueDate, err := time.Parse("2006-01-02", request.InvoiceIssueDate)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	dueDate, err := time.Parse("2006-01-02", request.InvoiceIssueDate)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	invoiceId := helper.GenerateInvoiceID(4)

	existsInvoice, _ := s.invoiceMysql.CheckExistsInvoiceId(invoiceId)
	if existsInvoice == true {
		err = fmt.Errorf("invoice_id '%s' already exists", invoiceId)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	_, err = s.customerMysql.FindById(request.Customer.IDCustomer)
	if err != nil {
		err = fmt.Errorf("customer_id '%s' not found", request.Customer.IDCustomer)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	customerIdUuid, err := uuid.Parse(request.Customer.IDCustomer)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	invoice = domain.Invoices{
		IDInvoice:         uuid.New(),
		CustomerID:        customerIdUuid,
		InvoiceID:         invoiceId,
		InvoiceSubject:    request.InvoiceSubject,
		InvoiceIssueDate:  issueDate,
		InvoiceDueDate:    dueDate,
		InvoiceTotalItem:  request.InvoiceTotalItem,
		InvoiceSubTotal:   request.InvoiceSubTotal,
		InvoiceTax:        request.InvoiceTax,
		InvoiceGrandTotal: request.InvoiceGrandTotal,
		InvoiceStatus:     request.InvoiceStatus,
		CreatedAt:         time.Now(),
	}

	for i, item := range request.Items {
		getItem, err = s.itemMysql.FindById(item.IDItem)
		if err != nil {
			err = fmt.Errorf("item_id.%d '%s' not found", i, request.Customer.IDCustomer)
			helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
			return "", err
		}

		if item.ItemQuantity > getItem.ItemQuantity {
			err = fmt.Errorf("item_id.%d with count %f is insufficient", i, item.ItemQuantity)
			helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
			return "", err
		}

		itemIdUuid, errUuid := uuid.Parse(item.IDItem)
		if errUuid != nil {
			helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{errUuid.Error()}, startTime)
			return "", errUuid
		}

		invoiceItems := domain.InvoiceHasItems{
			InvoiceID: invoice.IDInvoice,
			ItemID:    itemIdUuid,
			Quantity:  item.ItemQuantity,
			CreatedAt: time.Now(),
		}
		invoiceHasItems = append(invoiceHasItems, invoiceItems)

		newItemQuantity := getItem.ItemQuantity - item.ItemQuantity

		updateItem := domain.Items{
			IDItem:       itemIdUuid,
			ItemQuantity: newItemQuantity,
		}
		updateItems = append(updateItems, updateItem)
	}

	tx := s.invoiceMysql.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = tx.Create(&invoice).Error
	if err != nil {
		tx.Rollback()
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	for _, item := range invoiceHasItems {
		newItemQuantity := getItem.ItemQuantity - item.Quantity

		updateItem := domain.Items{
			IDItem:       item.ItemID,
			ItemQuantity: newItemQuantity,
		}

		updateItems = append(updateItems, updateItem)

		err = tx.Create(&item).Error
		if err != nil {
			tx.Rollback()
			helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
			return "", err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	for _, updateItem := range updateItems {
		err = s.itemMysql.Update(&updateItem)
		if err != nil {
			helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
			return "", err
		}
	}

	return helper.SuccessCreatedData, nil
}
