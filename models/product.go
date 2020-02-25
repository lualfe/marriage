package models

import "github.com/jinzhu/gorm"

// Product model
type Product struct {
	gorm.Model
	Room         string `gorm:"type:varchar(20)" json:"room"`
	Name         string `gorm:"type:varchar(20)" json:"name"`
	ProductModel string `gorm:"type:varchar(30)" json:"product_model"`
}
