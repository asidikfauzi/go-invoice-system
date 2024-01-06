package types

import (
	"go-invoice-system/model"
)

type TypesMysql interface {
	GetAll(limit, offset int, orderBy, typeName string) ([]model.Types, int64, error)
	FindById(typeId string) (model.Types, error)
}
