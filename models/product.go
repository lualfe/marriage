package models

import (
	"time"

	"github.com/google/uuid"
)

// Product model
type Product struct {
	ID           uuid.UUID  `gorm:"type:uuid" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	Room         string     `gorm:"type:varchar(20)" json:"room"`
	Name         string     `gorm:"type:varchar(20)" json:"name"`
	ProductModel string     `gorm:"type:varchar(30)" json:"product_model"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
