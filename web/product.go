package web

import (
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/models"
)

// FindProducts finds the products
func (w *Web) FindProducts(c echo.Context) error {
	p, r := w.App.Cockroach.FindProducts()
	r.JSON(c, p)
	return nil
}

// UpsertProduct upserts some products
func (w *Web) UpsertProduct(c echo.Context) error {
	p := &models.Product{}
	c.Bind(&p)
	p, r := w.App.Cockroach.UpsertProduct(p)
	r.JSON(c, p)
	return nil
}
