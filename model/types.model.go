package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Types struct {
		IDType    uuid.UUID  `json:"id_type"`
		TypeName  string     `json:"type_name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	GetType struct {
		IDType   uuid.UUID `json:"id_type"`
		TypeName string    `json:"type_name"`
	}

	RequestType struct {
		TypeName string `json:"type_name" validate:"required"`
	}
)
