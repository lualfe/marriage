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
	a.Instance.AutoMigrate(&models.Product{})
	a.Instance.AutoMigrate(&models.User{})
	a.Instance.AutoMigrate(&models.RecurringExpense{})
	a.Instance.AutoMigrate(&models.Couple{})
}

// BindForeignKeys binds foreign keys on the database
func (a *DB) BindForeignKeys() {
	fmt.Println("<-- Adding couple_id foreign key on users table -->")
	a.Instance.Model(&models.User{}).AddForeignKey("couple_id", "couples(id)", "SET NULL", "RESTRICT")
	fmt.Println("<-- Adding couple_id foreign key on products table -->")
	a.Instance.Model(&models.Product{}).AddForeignKey("couple_id", "couples(id)", "SET NULL", "RESTRICT")
	fmt.Println("<-- Adding couple_id foreign key on recurrent_expenses table -->")
	a.Instance.Model(&models.RecurringExpense{}).AddForeignKey("couple_id", "couples(id)", "SET NULL", "RESTRICT")
}
