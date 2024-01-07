package customers

import "go-invoice-system/model"

type CustomersMysql interface {
	GetAll(limit, offset int, orderBy, customerName string) ([]model.GetCustomer, int64, error)
	FindById(customerId string) (model.GetCustomer, error)
	FindByName(customerName string) (model.GetCustomer, error)
	CheckUpdateExists(customer model.Customers) (bool, error)
	Create(customer *model.Customers) error
	Update(customer *model.Customers) error
	Delete(customer *model.Customers) error
}
