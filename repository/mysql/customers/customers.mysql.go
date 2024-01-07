package customers

import "go-invoice-system/model"

type CustomersMysql interface {
	GetAll(limit, offset int, orderBy, typeName string) ([]model.GetCustomer, int64, error)
	FindById(typeId string) (model.GetCustomer, error)
	FindByName(typeName string) (model.GetCustomer, error)
	CheckUpdateExists(typ model.Customers) (bool, error)
	Create(typ *model.Customers) error
	Update(typ *model.Customers) error
	Delete(typ *model.Customers) error
}
