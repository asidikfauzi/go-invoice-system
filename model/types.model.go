package model

import (
	"github.com/google/uuid"
)

type (
	Types struct {
		IDType   uuid.UUID `json:"id_type"`
		TypeName string    `json:"type_name"`
	}

	RequestType struct {
		TypeName string `json:"type_name" validate:"required"`
	}
)
