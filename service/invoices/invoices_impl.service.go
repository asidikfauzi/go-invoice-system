package invoices

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"go-invoice-system/repository/mysql/invoices"
	"log"
	"math"
	"net/http"
	"time"
)

type Invoice struct {
	invoiceMysql invoices.InvoicesMysql
}

func NewInvoiceService(cm invoices.InvoicesMysql) InvoicesService {
	return &Invoice{
		invoiceMysql: cm,
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
