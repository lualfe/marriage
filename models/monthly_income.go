package models

import (
	"time"

	"github.com/google/uuid"
)

// MonthlyIncome model
type MonthlyIncome struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	CoupleID  uuid.UUID  `gorm:"type:uuid;UNIQUE;NOT NULL" json:"couple_id"`
	Value     float64    `gorm:"type:numeric; NOT NULL" json:"value" form:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
