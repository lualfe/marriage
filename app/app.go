package app

import (
	"github.com/jinzhu/gorm"
	"github.com/lualfe/casamento/app/cockroach"
	"github.com/spf13/viper"

	//needed for postgres initialization
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// Necessary in order to check for transaction retry error codes.
	_ "github.com/lib/pq"
)

// App instance
type App struct {
	Cockroach *cockroach.DB
}

//InitApp initializes an App instance
func InitApp() (*App, error) {
	db, err := gorm.Open("postgres", viper.GetString("COCKROACH_STRING"))
	if err != nil {
		panic("failed to connect database")
	}
	cr := &cockroach.DB{
		Instance: db,
	}
	return &App{Cockroach: cr}, nil
}
