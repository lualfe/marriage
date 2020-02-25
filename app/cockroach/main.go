package cockroach

import (
	"github.com/jinzhu/gorm"
	"github.com/lualfe/casamento/models"
)

// DB instance
type DB struct {
	Instance *gorm.DB
}

// Migrate migrate all the databases
func (a *DB) Migrate() {
	a.Instance.AutoMigrate(&models.Product{})
}
