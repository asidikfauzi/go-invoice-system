package customers

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/common/validator"
	"go-invoice-system/model"
	"log"
	"net/http"
	"time"
)

func (m *MasterCustomers) GetAllCustomers(c *gin.Context) {
	startTime := time.Now()

	var request model.Customers
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	pageParam := c.Query("page")
	limitParam := c.Query("limit")
	orderByParam := c.Query("orderBy")

	dataCustomer, paginate, err := m.CustomerService.GetAllCustomers(c, pageParam, limitParam, orderByParam, request.CustomerName, startTime)
	if err != nil {
		log.Printf("error customer controller GetAllCustomers :%s", err)
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataCustomer, paginate, startTime)
	return
}

func (m *MasterCustomers) FindCustomerById(c *gin.Context) {
	startTime := time.Now()

	customerId := c.Param("customerId")
	dataCustomer, err := m.CustomerService.FindCustomerById(c, customerId, startTime)
	if err != nil {
		log.Printf("error customer controller FindCustomerById :%s", err)
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, helper.Success, []string{helper.SuccessGetData}, dataCustomer, startTime)
	return
}

func (m *MasterCustomers) CreateCustomer(c *gin.Context) {
	startTime := time.Now()

	var request model.RequestCustomer
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.CustomerService.CreateCustomer(c, request, startTime)
	if err != nil {
		log.Printf("error customer controller CreateCustomer :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, helper.SuccessCreated, []string{msg}, startTime)
	return

}

func (m *MasterCustomers) UpdateCustomer(c *gin.Context) {
	startTime := time.Now()

	customerId := c.Param("customerId")
	var request model.RequestCustomer
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(request)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, validate, startTime)
		return
	}

	msg, err := m.CustomerService.UpdateCustomer(c, request, customerId, startTime)
	if err != nil {
		log.Printf("error customer controller UpdateCustomer :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, helper.Success, []string{msg}, startTime)
	return

}

func (m *MasterCustomers) DeleteCustomer(c *gin.Context) {
	startTime := time.Now()

	customerId := c.Param("customerId")
	msg, err := m.CustomerService.DeleteCustomer(c, customerId, startTime)
	if err != nil {
		log.Printf("error customer controller DeleteCustomer :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, helper.Success, []string{msg}, startTime)
	return

}
