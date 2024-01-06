package domain

import (
	"github.com/google/uuid"
	"time"
)

type InvoiceHasItems struct {
	InvoiceID   uuid.UUID  `gorm:"type:char(36);not null" json:"invoice_id"`
	ItemID      uuid.UUID  `gorm:"type:char(36);not null" json:"item_id"`
	Quantity    float64    `gorm:"type:float;" json:"quantity"`
	CreatedAt   time.Time  `gorm:"type:timestamp;default:current_timestamp()" json:"created_at"`
	CreatedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"created_by_id"`
	UpdatedAt   *time.Time `gorm:"default:null" json:"updated_at"`
	UpdatedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"updated_by_id"`
	DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"deleted_by_id"`

	// REFERENCES
	Invoice Invoices `gorm:"foreignKey:InvoiceID;references:id_invoice" json:"invoice"`
	Item    Items    `gorm:"foreignKey:ItemID;references:id_item" json:"item"`
}
