package customers

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/service/customers"
)

type CustomerController interface {
	GetAllCustomers(c *gin.Context)
	FindCustomerById(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type MasterCustomers struct {
	CustomerService customers.CustomersService `inject:"customers_service"`
}
