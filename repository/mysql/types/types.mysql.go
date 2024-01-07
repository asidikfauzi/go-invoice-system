package types

import (
	"go-invoice-system/model"
)

type TypesMysql interface {
	GetAll(limit, offset int, orderBy, typeName string) ([]model.GetType, int64, error)
	FindById(typeId string) (model.GetType, error)
	FindByName(typeName string) (model.GetType, error)
	CheckUpdateExists(typ model.Types) (bool, error)
	Create(typ *model.Types) error
	Update(typ *model.Types) error
	Delete(typ *model.Types) error
}
