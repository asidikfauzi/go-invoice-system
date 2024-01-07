package items

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
)

type ItemsMysql interface {
	GetAll(limit, offset int, orderBy, itemName string) ([]model.GetItem, int64, error)
	FindById(itemId string) (model.GetItem, error)
	FindByName(itemName string) (model.GetItem, error)
	CheckUpdateExists(typ domain.Items) (bool, error)
	Create(item *domain.Items) error
	Update(item *domain.Items) error
	Delete(item *domain.Items) error
}
