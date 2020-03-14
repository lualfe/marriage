package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lualfe/casamento/utils"
	"github.com/spf13/viper"
)

// InitRoutes initializes all routes
func (w *Web) InitRoutes(e *echo.Echo) {
	// Routes that does not require auth
	e.POST("/user/register", w.CreateUser, utils.CheckToken)
	e.POST("/user/login", w.LoginUser, utils.CheckToken)
	e.PUT("/user/update/:id", w.UpdateUser)

	// Routes that require auth
	userAuthRequired := e.Group("/user")
	userAuthRequired.Use(middleware.JWT([]byte(viper.GetString("jwt_key"))))
	userAuthRequired.GET("/couple/invite", w.PartnerInvite)
	userAuthRequired.GET("/products", w.FindProducts)
	userAuthRequired.POST("/product", w.UpsertProduct)
}
