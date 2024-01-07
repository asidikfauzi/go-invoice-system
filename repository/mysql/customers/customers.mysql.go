package customers

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
)

type CustomersMysql interface {
	GetAll(limit, offset int, orderBy, customerName string) ([]model.GetCustomer, int64, error)
	FindById(customerId string) (model.GetCustomer, error)
	FindByName(customerName string) (model.GetCustomer, error)
	CheckUpdateExists(customer domain.Customers) (bool, error)
	Create(customer *domain.Customers) error
	Update(customer *domain.Customers) error
	Delete(customer *domain.Customers) error
}
