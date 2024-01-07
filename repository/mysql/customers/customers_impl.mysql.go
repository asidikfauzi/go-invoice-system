package customers

import (
	"go-invoice-system/model"
	"go-invoice-system/model/domain"
	"gorm.io/gorm"
)

type Customers struct {
	DB *gorm.DB
}

func NewCustomerMysql(conn *gorm.DB) CustomersMysql {
	return &Customers{
		DB: conn,
	}
}

func (m *Customers) GetAll(limit, offset int, orderBy, customerName string) ([]model.GetCustomer, int64, error) {
	var (
		customers  []domain.Customers
		data       []model.GetCustomer
		totalCount int64
	)

	query := m.DB.Where("deleted_at IS NULL")

	if customerName != "" {
		query = query.Where("customer_name LIKE ?", "%"+customerName+"%")
	}

	if orderBy != "" {
		query = query.Order("customer_name " + orderBy)
	} else {
		query = query.Order("customer_name ASC")
	}

	if err := query.Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, totalCount, err
	}

	if err := query.Model(&domain.Customers{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	for _, c := range customers {
		customer := model.GetCustomer{
			IDCustomer:      c.IDCustomer,
			CustomerName:    c.CustomerName,
			CustomerAddress: c.CustomerAddress,
		}
		data = append(data, customer)
	}

	return data, totalCount, nil
}

func (m *Customers) FindById(customerId string) (model.GetCustomer, error) {
	var (
		customer domain.Customers
		data     model.GetCustomer
	)

	if err := m.DB.Where("id_customer = ?", customerId).
		Where("deleted_at IS NULL").
		First(&customer).Error; err != nil {
		return data, err
	}

	data = model.GetCustomer{
		IDCustomer:      customer.IDCustomer,
		CustomerName:    customer.CustomerName,
		CustomerAddress: customer.CustomerAddress,
	}

	return data, nil
}

func (m *Customers) FindByName(customerName string) (model.GetCustomer, error) {
	var (
		customer domain.Customers
		data     model.GetCustomer
	)

	if err := m.DB.Where("customer_name = ?", customerName).First(&customer).Error; err != nil {
		return data, err
	}

	data = model.GetCustomer{
		IDCustomer:      customer.IDCustomer,
		CustomerName:    customer.CustomerName,
		CustomerAddress: customer.CustomerAddress,
	}

	return data, nil
}

func (m *Customers) CheckUpdateExists(customer domain.Customers) (bool, error) {
	err := m.DB.Where("customer_name = ?", customer.CustomerName).
		Where("id_customer != ? ", customer.IDCustomer).
		Where("deleted_at IS NULL").First(&customer).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Customers) Create(customer *domain.Customers) error {
	return m.DB.Create(customer).Error
}

func (m *Customers) Update(customer *domain.Customers) error {
	updateCustomer := domain.Customers{
		CustomerName:    customer.CustomerName,
		CustomerAddress: customer.CustomerAddress,
		UpdatedAt:       customer.UpdatedAt,
	}
	return m.DB.Where("id_customer = ?", customer.IDCustomer).Updates(updateCustomer).Error
}

func (m *Customers) Delete(customer *domain.Customers) error {
	return m.DB.Model(&domain.Customers{}).Where("id_customer = ?", customer.IDCustomer).
		UpdateColumn("deleted_at", customer.DeletedAt).Error
}
