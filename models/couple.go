package models

import (
	"time"

	"github.com/google/uuid"
)

// Couple model
type Couple struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	Token     string     `gorm:"type:varchar(10)" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
