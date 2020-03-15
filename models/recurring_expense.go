package models

import (
	"time"

	"github.com/google/uuid"
)

// RecurringExpense model
type RecurringExpense struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	CoupleID  uuid.UUID  `gorm:"type:uuid;NOT NULL" json:"couple_id"`
	Name      string     `gorm:"type:varchar(50)" json:"name" form:"name"`
	Value     float64    `gorm:"type:numeric; NOT NULL" json:"value" form:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
