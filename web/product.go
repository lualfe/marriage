package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/app/cockroach"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// FindProducts finds the products
func (w *Web) FindProducts(c echo.Context) error {
	jwt := utils.JWTGetter(c, "id")
	id := jwt[0].(string)

	// set list of filters
	filters := make([]string, 0)
	room := c.QueryParam("room")
	if room != "" {
		filters = append(filters, "room")
	}
	brand := c.QueryParam("brand")
	if brand != "" {
		filters = append(filters, "brand")
	}
	name := c.QueryParam("name")
	if name != "" {
		filters = append(filters, "name")
	}

	// set product finder interface
	var finder cockroach.ProductFinder
	switch len(filters) {
	case 0:
		finder = &cockroach.IDFinder{
			ID: id,
		}
	case 1:
		switch filters[0] {
		case "room":
			finder = &cockroach.RoomFinder{
				UserID: id,
				Room:   room,
			}
		case "brand":
			brand := c.QueryParam("brand")
			finder = &cockroach.BrandFinder{
				UserID: id,
				Brand:  brand,
			}
		case "name":
			name := c.QueryParam("name")
			finder = &cockroach.NameFinder{
				UserID: id,
				Name:   name,
			}
		}
	default:
		multipleFinder := &cockroach.MultipleFinder{}
		for _, filter := range filters {
			switch filter {
			case "id":
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.IDFinder{
					ID: id,
				})
			case "room":
				room := c.QueryParam("room")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.RoomFinder{
					UserID: id,
					Room:   room,
				})
			case "brand":
				brand := c.QueryParam("brand")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.BrandFinder{
					UserID: id,
					Brand:  brand,
				})
			case "name":
				name := c.QueryParam("name")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.NameFinder{
					UserID: id,
					Name:   name,
				})
			}
		}
		finder = multipleFinder
	}
	p, r := w.App.Cockroach.FindProducts(finder)
	r.JSON(c, p)
	return nil
}

// UpsertProduct upserts some products
func (w *Web) UpsertProduct(c echo.Context) error {
	jwt := utils.JWTGetter(c, "id")
	userID := jwt[0].(string)
	p := &models.Product{}
	c.Bind(&p)
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	userIDAsUUID, err := uuid.Parse(userID)
	if err != nil {
		response := responsewriter.UnexpectedError(err)
		response.JSON(c, "")
		return nil
	}
	p.UserID = userIDAsUUID
	product, r := w.App.Cockroach.UpsertProduct(p)
	r.JSON(c, product)
	return nil
}
