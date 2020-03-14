package models

import (
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID                uuid.UUID  `gorm:"type:uuid" json:"id"`
	CoupleID          uuid.UUID  `gorm:"type:uuid" json:"couple_id"`
	Name              string     `gorm:"type:varchar(20); NOT NULL" json:"name"`
	Email             string     `gorm:"type:varchar(60);NOT NULL" json:"email"`
	Password          string     `gorm:"type:varchar(200);NOT NULL" json:"-"`
	TemporaryPassword string     `gorm:"type:varchar(200); NULL" json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
