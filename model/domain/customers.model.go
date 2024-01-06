package domain

import (
	"github.com/google/uuid"
	"time"
)

type Customers struct {
	IDCustomer      uuid.UUID  `gorm:"type:char(36);unique;not null;primary_key;column:id_customer" json:"id_customer"`
	CustomerName    string     `gorm:"type:varchar(50);not null;" json:"customer_name"`
	CustomerAddress string     `gorm:"type:text;" json:"customer_address"`
	CreatedAt       time.Time  `gorm:"type:timestamp;default:current_timestamp()" json:"created_at"`
	CreatedByID     *uuid.UUID `gorm:"type:char(36);default:null" json:"created_by_id"`
	UpdatedAt       *time.Time `gorm:"default:null" json:"updated_at"`
	UpdatedByID     *uuid.UUID `gorm:"type:char(36);default:null" json:"updated_by_id"`
	DeletedAt       *time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedByID     *uuid.UUID `gorm:"type:char(36);default:null" json:"deleted_by_id"`
}
