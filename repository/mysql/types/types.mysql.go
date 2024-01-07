package types

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
)

type TypesMysql interface {
	GetAll(limit, offset int, orderBy, typeName string) ([]model.GetType, int64, error)
	FindById(typeId string) (model.GetType, error)
	FindByName(typeName string) (model.GetType, error)
	CheckUpdateExists(typ domain.Types) (bool, error)
	Create(typ *domain.Types) error
	Update(typ *domain.Types) error
	Delete(typ *domain.Types) error
}
