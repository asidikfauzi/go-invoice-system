package items

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"gorm.io/gorm"
)

type Items struct {
	DB *gorm.DB
}

func NewItemMysql(conn *gorm.DB) ItemsMysql {
	return &Items{
		DB: conn,
	}
}

func (m *Items) GetAll(limit, offset int, orderBy, itemName string) ([]model.GetItem, int64, error) {
	var (
		items      []domain.Items
		data       []model.GetItem
		totalCount int64
	)

	query := m.DB.Where("items.deleted_at IS NULL")

	if itemName != "" {
		query = query.Where("items.item_name LIKE ?", "%"+itemName+"%")
	}

	if orderBy != "" {
		query = query.Order("item_name " + orderBy)
	} else {
		query = query.Order("item_name ASC")
	}

	if err := query.Preload("Types").
		Joins("JOIN types t ON items.type_id = t.id_type").
		Offset(offset).
		Limit(limit).
		Find(&items).Error; err != nil {
		return nil, totalCount, err
	}

	if err := query.Model(&domain.Items{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	for _, i := range items {
		item := model.GetItem{
			IDItem:       i.IDItem,
			ItemName:     i.ItemName,
			ItemQuantity: i.ItemQuantity,
			ItemPrice:    i.ItemPrice,
			TypeID:       i.Types.IDType,
			TypeName:     i.Types.TypeName,
		}
		data = append(data, item)
	}

	return data, totalCount, nil
}

func (m *Items) FindById(itemId string) (model.GetItem, error) {
	var (
		item domain.Items
		data model.GetItem
	)

	if err := m.DB.Preload("Types").
		Joins("JOIN types t ON items.type_id = t.id_type").
		Where("id_item = ?", itemId).
		Where("items.deleted_at IS NULL").
		First(&item).Error; err != nil {
		return data, err
	}

	data = model.GetItem{
		IDItem:       item.IDItem,
		ItemName:     item.ItemName,
		ItemQuantity: item.ItemQuantity,
		ItemPrice:    item.ItemPrice,
		TypeID:       item.Types.IDType,
		TypeName:     item.Types.TypeName,
	}

	return data, nil
}

func (m *Items) FindByName(itemName string) (model.GetItem, error) {
	var (
		item domain.Items
		data model.GetItem
	)

	if err := m.DB.Preload("Types").
		Joins("JOIN types t ON items.type_id = t.id_type").
		Where("items.item_name = ?", itemName).
		First(&item).Error; err != nil {
		return data, err
	}

	data = model.GetItem{
		IDItem:       item.IDItem,
		ItemName:     item.ItemName,
		ItemQuantity: item.ItemQuantity,
		ItemPrice:    item.ItemPrice,
		TypeID:       item.Types.IDType,
		TypeName:     item.Types.TypeName,
	}

	return data, nil
}

func (m *Items) CheckUpdateExists(item domain.Items) (bool, error) {
	var (
		count  int64
		status bool
	)

	err := m.DB.Model(&domain.Items{}).
		Where("item_name = ?", item.ItemName).
		Where("id_item != ? ", item.IDItem).
		Where("deleted_at IS NULL").
		Count(&count).Error

	status = false
	if count > 0 {
		status = true
	}

	if err != nil {
		return status, err
	}

	return status, nil
}

func (m *Items) Create(item *domain.Items) error {
	return m.DB.Create(item).Error
}

func (m *Items) Update(item *domain.Items) error {
	updateItem := domain.Items{
		ItemName:     item.ItemName,
		ItemQuantity: item.ItemQuantity,
		ItemPrice:    item.ItemPrice,
		TypeID:       item.Types.IDType,
	}
	return m.DB.Where("id_item = ?", item.IDItem).Updates(updateItem).Error
}

func (m *Items) Delete(item *domain.Items) error {
	return m.DB.Model(&domain.Items{}).Where("id_item = ?", item.IDItem).
		UpdateColumn("deleted_at", item.DeletedAt).Error
}
