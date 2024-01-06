package types

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
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

func (m *Types) FindById(typeId string) (model.Types, error) {
	var typ model.Types

	if err := m.DB.Where("id_type = ?", typeId).First(&typ).Error; err != nil {
		return typ, err
	}

	return typ, nil
}

func (m *Types) FindByName(typeName string) (model.Types, error) {
	var typ model.Types

	if err := m.DB.Where("type_name = ?", typeName).First(&typ).Error; err != nil {
		return typ, err
	}

	return typ, nil
}

func (m *Types) Create(typ *domain.Types) error {
	return m.DB.Create(typ).Error
}
