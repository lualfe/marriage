package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// FindMonthlyIncome finds a couple income
func (w *Web) FindMonthlyIncome(c echo.Context) error {
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)
	coupleID, _ := uuid.Parse(cID)
	income, response := w.App.Cockroach.FindMonthlyIncome(coupleID)
	response.JSON(c, income)
	return nil
}

// UpsertMonthlyIncome upserts an couple income
func (w *Web) UpsertMonthlyIncome(c echo.Context) error {
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)

	coupleID, _ := uuid.Parse(cID)
	income, response := w.App.Cockroach.FindMonthlyIncome(coupleID)
	if income == nil {
		income = &models.MonthlyIncome{}
	}
	c.Bind(income)
	income.CoupleID = coupleID

	if income.ID == uuid.Nil {
		income.ID = uuid.New()
	}
	income, response = w.App.Cockroach.UpsertMonthlyIncome(income)
	response.JSON(c, income)
	return nil
}
