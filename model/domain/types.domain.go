package domain

import (
	"github.com/google/uuid"
	"time"
)

type Types struct {
	IDType      uuid.UUID  `gorm:"type:char(36);unique;not null;primary_key;column:id_type" json:"id_type"`
	TypeName    string     `gorm:"type:varchar(50);not null;" json:"type_name"`
	CreatedAt   time.Time  `gorm:"default:null" json:"created_at"`
	CreatedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"created_by_id"`
	UpdatedAt   *time.Time `gorm:"default:null" json:"updated_at"`
	UpdatedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"updated_by_id"`
	DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedByID *uuid.UUID `gorm:"type:char(36);default:null" json:"deleted_by_id"`
}
