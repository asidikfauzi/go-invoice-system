package invoices

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"log"
	"net/http"
	"time"
)

func (m *MasterInvoices) GetAllInvoices(c *gin.Context) {
	startTime := time.Now()

	var request model.RequestInvoices
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	pageParam := c.Query("page")
	limitParam := c.Query("limit")
	orderByParam := c.Query("orderBy")

	dataInvoice, paginate, err := m.InvoiceService.GetAllInvoices(c, pageParam, limitParam, orderByParam, request, startTime)
	if err != nil {
		log.Printf("error invoice controller GetAllItems :%s", err)
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataInvoice, paginate, startTime)
	return
}
