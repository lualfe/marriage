package models

import (
	"time"

	"github.com/google/uuid"
)

// Product model
type Product struct {
	ID           uuid.UUID  `gorm:"type:uuid" json:"id"`
	CoupleID     uuid.UUID  `gorm:"type:uuid;NOT NULL" json:"user_id"`
	Room         string     `gorm:"type:varchar(20);NOT NULL" json:"room"`
	Name         string     `gorm:"type:varchar(20); NOT NULL" json:"name"`
	Brand        string     `gorm:"varchar(10); NOT NULL" json:"brand"`
	ProductModel string     `gorm:"type:varchar(30); NOT NULL" json:"product_model"`
	Installments int        `gorm:"type:int; NULL" json:"installments"`
	Gift         bool       `gorm:"type:bool;DEFAULT:false" json:"gift"`
	Price        float64    `gorm:"type:numeric;NOT NULL" json:"price"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
