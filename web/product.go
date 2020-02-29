package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// FindProducts finds the products
func (w *Web) FindProducts(c echo.Context) error {
	jwt := utils.JWTGetter(c, "id")
	id := jwt[0].(string)
	p, r := w.App.Cockroach.FindProducts(id)
	r.JSON(c, p)
	return nil
}

// UpsertProduct upserts some products
func (w *Web) UpsertProduct(c echo.Context) error {
	p := &models.Product{}
	c.Bind(&p)
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	p, r := w.App.Cockroach.UpsertProduct(p)
	r.JSON(c, p)
	return nil
}
