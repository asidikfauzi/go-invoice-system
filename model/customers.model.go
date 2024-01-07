package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Customers struct {
		IDCustomer      uuid.UUID  `json:"id_customer"`
		CustomerName    string     `json:"customer_name"`
		CustomerAddress string     `json:"customer_address"`
		CreatedAt       time.Time  `json:"created_at"`
		UpdatedAt       *time.Time `json:"updated_at"`
		DeletedAt       *time.Time `json:"deleted_at"`
	}

	GetCustomer struct {
		IDCustomer      uuid.UUID `json:"id_customer"`
		CustomerName    string    `json:"customer_name"`
		CustomerAddress string    `json:"customer_address"`
	}

	RequestCustomer struct {
		CustomerName    string `json:"customer_name" validate:"required"`
		CustomerAddress string `json:"customer_address" validate:"required"`
	}
)
