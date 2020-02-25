package web

import (
	"github.com/labstack/echo"
)

// InitRoutes initializes all routes
func (w *Web) InitRoutes(e *echo.Echo) {
	e.GET("/", w.FindProducts)
}
