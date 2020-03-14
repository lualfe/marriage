package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/app/cockroach"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// FindProducts finds the products
func (w *Web) FindProducts(c echo.Context) error {
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)
	coupleID, _ := uuid.Parse(cID)

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
		finder = &cockroach.CoupleIDFinder{
			CoupleID: coupleID,
		}
	case 1:
		switch filters[0] {
		case "room":
			finder = &cockroach.RoomFinder{
				CoupleID: coupleID,
				Room:     room,
			}
		case "brand":
			brand := c.QueryParam("brand")
			finder = &cockroach.BrandFinder{
				CoupleID: coupleID,
				Brand:    brand,
			}
		case "name":
			name := c.QueryParam("name")
			finder = &cockroach.NameFinder{
				CoupleID: coupleID,
				Name:     name,
			}
		}
	default:
		multipleFinder := &cockroach.MultipleFinder{}
		for _, filter := range filters {
			switch filter {
			case "user_id":
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.CoupleIDFinder{
					CoupleID: coupleID,
				})
			case "room":
				room := c.QueryParam("room")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.RoomFinder{
					CoupleID: coupleID,
					Room:     room,
				})
			case "brand":
				brand := c.QueryParam("brand")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.BrandFinder{
					CoupleID: coupleID,
					Brand:    brand,
				})
			case "name":
				name := c.QueryParam("name")
				multipleFinder.Multiple = append(multipleFinder.Multiple, &cockroach.NameFinder{
					CoupleID: coupleID,
					Name:     name,
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
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)
	p := &models.Product{}
	c.Bind(&p)
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	coupleID, _ := uuid.Parse(cID)
	p.CoupleID = coupleID
	product, r := w.App.Cockroach.UpsertProduct(p)
	r.JSON(c, product)
	return nil
}
