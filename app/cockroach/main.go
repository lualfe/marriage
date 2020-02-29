package cockroach

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lualfe/casamento/models"
)

// DB instance
type DB struct {
	Instance *gorm.DB
}

// Migrate migrate all the databases
func (a *DB) Migrate() {
	go func() {
		a.Instance.AutoMigrate(&models.Product{})
		a.Instance.AutoMigrate(&models.User{})
	}()
}

// BindForeignKeys binds foreign keys on the database
func (a *DB) BindForeignKeys() {
	go func() {
		a.Instance.LogMode(false)
		fmt.Println("<-- Adding user_id foreign key on products table -->")
		a.Instance.Model(&models.Product{}).AddForeignKey("user_id", "users(id)", "NO ACTION", "RESTRICT")
		a.Instance.LogMode(true)
	}()
}
