package models

import (
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID        uuid.UUID  `gorm:"type:uuid" json:"id"`
	Name      string     `gorm:"type:varchar(20)" json:"name"`
	LastName  string     `gorm:"type:varchar(20)" json:"last_name"`
	Email     string     `gorm:"type:varchar(60);NOT NULL" json:"email"`
	Password  string     `gorm:"type:varchar(200);NOT NULL" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
