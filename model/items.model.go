package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Items struct {
		IDItem       uuid.UUID  `json:"id_item"`
		ItemName     string     `json:"item_name"`
		ItemQuantity float64    `json:"item_quantity"`
		ItemPrice    float64    `json:"item_price"`
		TypeID       uuid.UUID  `json:"type_id"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at"`
		DeletedAt    *time.Time `json:"deleted_at"`
		Types        Types      `gorm:"foreignKey:TypeID" json:"types"`
	}

	GetItem struct {
		IDItem       uuid.UUID `json:"id_item"`
		ItemName     string    `json:"item_name"`
		ItemQuantity float64   `json:"item_quantity"`
		ItemPrice    float64   `json:"item_price"`
		TypeID       uuid.UUID `json:"type_id"`
		TypeName     string    `json:"type_name"`
	}

	RequestItem struct {
		ItemName     string  `json:"item_name" validate:"required"`
		ItemQuantity float64 `json:"item_quantity" validate:"required,number"`
		ItemPrice    float64 `json:"item_price" validate:"required,number"`
		TypeID       string  `json:"type_id" validate:"required"`
	}
)
