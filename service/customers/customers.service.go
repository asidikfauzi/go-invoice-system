package customers

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/model"
	"time"
)

type CustomersService interface {
	GetAllCustomers(c *gin.Context, pageParam, limitParam, orderByParam, customerName string, startTime time.Time) ([]model.GetCustomer, helper.Paginate, error)
	FindCustomerById(c *gin.Context, customerId string, startTime time.Time) (model.GetCustomer, error)
	CreateCustomer(c *gin.Context, request model.RequestCustomer, startTime time.Time) (string, error)
	UpdateCustomer(c *gin.Context, request model.RequestCustomer, customerId string, startTime time.Time) (string, error)
	DeleteCustomer(c *gin.Context, customerId string, startTime time.Time) (string, error)
}
