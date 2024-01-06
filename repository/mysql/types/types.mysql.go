package types

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
)

type TypesMysql interface {
	GetAll(limit, offset int, orderBy, typeName string) ([]model.Types, int64, error)
	FindById(typeId string) (model.Types, error)
	FindByName(typeName string) (model.Types, error)
	Create(typ *domain.Types) error
}
