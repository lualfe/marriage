package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lualfe/casamento/app"
	"github.com/lualfe/casamento/web"
	"github.com/spf13/viper"
)

func newEcho() *echo.Echo {
	return echo.New()
}

func configureVariables() {
	viper.SetConfigName("default")   // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	configureVariables()
	e := newEcho()
	a, err := app.InitApp()
	defer a.Cockroach.Instance.Close()
	if err != nil {
		e.Logger.Printf("error initializing database: %v", err)
		panic(err)
	}
	a.Cockroach.Migrate()
	w, err := web.New(a)
	if err != nil {
		e.Logger.Printf("error initializing database: %v", err)
		panic(err)
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	w.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
