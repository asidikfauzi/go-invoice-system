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

func (m *Types) GetAll(limit, offset int, orderBy, typeName string) ([]model.GetType, int64, error) {
	var (
		types      []domain.Types
		data       []model.GetType
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

	if err := query.Model(&domain.Types{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	for _, t := range types {
		typ := model.GetType{
			IDType:   t.IDType,
			TypeName: t.TypeName,
		}
		data = append(data, typ)
	}

	return data, totalCount, nil
}

func (m *Types) FindById(typeId string) (model.GetType, error) {
	var (
		typ  domain.Types
		data model.GetType
	)

	if err := m.DB.Where("id_type = ?", typeId).
		Where("deleted_at IS NULL").
		First(&typ).Error; err != nil {
		return data, err
	}

	data = model.GetType{
		IDType:   typ.IDType,
		TypeName: typ.TypeName,
	}

	return data, nil
}

func (m *Types) FindByName(typeName string) (model.GetType, error) {
	var (
		typ  domain.Types
		data model.GetType
	)

	if err := m.DB.Where("type_name = ?", typeName).First(&typ).Error; err != nil {
		return data, err
	}

	data = model.GetType{
		IDType:   typ.IDType,
		TypeName: typ.TypeName,
	}

	return data, nil
}

func (m *Types) CheckUpdateExists(typ domain.Types) (bool, error) {
	err := m.DB.Where("type_name = ?", typ.TypeName).
		Where("id_type != ? ", typ.IDType).
		Where("deleted_at IS NULL").First(&typ).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Types) Create(typ *domain.Types) error {
	return m.DB.Create(typ).Error
}

func (m *Types) Update(typ *domain.Types) error {
	updateType := domain.Types{
		TypeName:  typ.TypeName,
		UpdatedAt: typ.UpdatedAt,
	}
	return m.DB.Where("id_type = ?", typ.IDType).Updates(updateType).Error
}

func (m *Types) Delete(typ *domain.Types) error {
	return m.DB.Model(&domain.Types{}).Where("id_type = ?", typ.IDType).
		UpdateColumn("deleted_at", typ.DeletedAt).Error
}
