package types

import (
	"go-invoice-system/model"
	"gorm.io/gorm"
)

type Types struct {
	DB *gorm.DB
}

func NewTypeMysql(conn *gorm.DB) TypesMysql {
	return &Types{
		DB: conn,
	}
}

func (m *Types) GetAll(limit, offset int, orderBy, typeName string) ([]model.Types, int64, error) {
	var (
		types      []model.Types
		totalCount int64
	)

	query := m.DB.Where("deleted_at IS NULL")

	if typeName != "" {
		query = query.Where("type_name LIKE ?", "%"+typeName+"%")
	}

	if orderBy != "" {
		query = query.Order("type_name " + orderBy)
	} else {
		query = query.Order("type_name ASC")
	}

	if err := query.Offset(offset).Limit(limit).Find(&types).Error; err != nil {
		return nil, totalCount, err
	}

	if err := query.Model(&model.Types{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return types, totalCount, nil
}
