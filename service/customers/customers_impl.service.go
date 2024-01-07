package customers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"go-invoice-system/repository/mysql/customers"
	"log"
	"math"
	"net/http"
	"time"
)

type Customer struct {
	customerMysql customers.CustomersMysql
}

func NewCustomerService(cm customers.CustomersMysql) CustomersService {
	return &Customer{
		customerMysql: cm,
	}
}

func (s *Customer) GetAllCustomers(c *gin.Context, pageParam, limitParam, orderByParam, customerName string, startTime time.Time) ([]model.GetCustomer, helper.Paginate, error) {
	var (
		dataCustomers []model.GetCustomer
		paginate      helper.Paginate
		totalData     int64
		err           error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return dataCustomers, paginate, err
	}

	dataCustomers, totalData, err = s.customerMysql.GetAll(limit, offset, orderByParam, customerName)
	if err != nil {
		log.Printf("error customer service GetAll: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return dataCustomers, paginate, err
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	paginate = helper.Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return dataCustomers, paginate, nil
}

func (s *Customer) FindCustomerById(c *gin.Context, customerId string, startTime time.Time) (model.GetCustomer, error) {
	var (
		customer model.GetCustomer
		err      error
	)

	customer, err = s.customerMysql.FindById(customerId)
	if err != nil {
		err = fmt.Errorf("customer_id '%s' not found", customerId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return customer, err
	}

	return customer, nil
}

func (s *Customer) CreateCustomer(c *gin.Context, request model.RequestCustomer, startTime time.Time) (string, error) {
	var err error

	findByName, _ := s.customerMysql.FindByName(request.CustomerName)

	if findByName.CustomerName != "" {
		err = fmt.Errorf("customer_name '%s' already exists", request.CustomerName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	customers := domain.Customers{
		IDCustomer:      uuid.New(),
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		CreatedAt:       time.Now(),
	}

	err = s.customerMysql.Create(&customers)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessCreatedData, nil
}

func (s *Customer) UpdateCustomer(c *gin.Context, request model.RequestCustomer, customerId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.customerMysql.FindById(customerId)
	if err != nil {
		err = fmt.Errorf("customer_id '%s' not found", customerId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	customerIdUuid, err := uuid.Parse(customerId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeUpdate := time.Now()
	customers := domain.Customers{
		IDCustomer:      customerIdUuid,
		CustomerName:    request.CustomerName,
		CustomerAddress: request.CustomerAddress,
		UpdatedAt:       &timeUpdate,
	}

	exists, _ := s.customerMysql.CheckUpdateExists(customers)
	if exists == true {
		err = fmt.Errorf("customer_name '%s' already exists", request.CustomerName)
		helper.ResponseAPI(c, false, http.StatusConflict, helper.Conflict, []string{err.Error()}, startTime)
		return "", err
	}

	err = s.customerMysql.Update(&customers)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessUpdatedData, nil
}

func (s *Customer) DeleteCustomer(c *gin.Context, customerId string, startTime time.Time) (string, error) {
	var err error

	_, err = s.customerMysql.FindById(customerId)
	if err != nil {
		err = fmt.Errorf("customer_id '%s' not found", customerId)
		helper.ResponseAPI(c, false, http.StatusNotFound, helper.NotFound, []string{err.Error()}, startTime)
		return "", err
	}

	customerIdUuid, err := uuid.Parse(customerId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, helper.BadRequest, []string{err.Error()}, startTime)
		return "", err
	}

	timeDelete := time.Now()
	customers := domain.Customers{
		IDCustomer: customerIdUuid,
		DeletedAt:  &timeDelete,
	}

	err = s.customerMysql.Delete(&customers)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, helper.InternalServerError, []string{err.Error()}, startTime)
		return "", err
	}

	return helper.SuccessDeletedData, nil
}
