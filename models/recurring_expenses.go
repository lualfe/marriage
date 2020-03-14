package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// RecurringExpense model
type RecurringExpense struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	CoupleID  uuid.UUID  `gorm:"type:uuid" json:"couple_id"`
	Value     float64    `gorm:"type:numeric; NOT NULL" json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
