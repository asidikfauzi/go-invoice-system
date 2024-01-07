package domain

import (
	"github.com/google/uuid"
	"time"
)

type Items struct {
	IDItem       uuid.UUID  `gorm:"type:char(36);unique;not null;primary_key;column:id_item" json:"id_item"`
	TypeID       uuid.UUID  `gorm:"type:char(36);not null;" json:"type_id"`
	ItemName     string     `gorm:"type:varchar(125)" json:"item_name"`
	ItemQuantity float64    `gorm:"type:float" json:"item_quantity"`
	ItemPrice    float64    `gorm:"type:float" json:"item_price"`
	CreatedAt    time.Time  `gorm:"default:null" json:"created_at"`
	CreatedByID  *uuid.UUID `gorm:"type:char(36);default:null" json:"created_by_id"`
	UpdatedAt    *time.Time `gorm:"default:null" json:"updated_at"`
	UpdatedByID  *uuid.UUID `gorm:"type:char(36);default:null" json:"updated_by_id"`
	DeletedAt    *time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedByID  *uuid.UUID `gorm:"type:char(36);default:null" json:"deleted_by_id"`

	//REFERENCE
	Types Types `gorm:"foreignKey:TypeID;references:id_type" json:"types"`
}
