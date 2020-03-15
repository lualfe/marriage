package web

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// UpsertExpense inserts an expense in database. If already existis, updates
func (w *Web) UpsertExpense(c echo.Context) error {
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)

	expense := &models.RecurringExpense{}
	c.Bind(expense)
	if c.Param("id") != "" {
		expenseID, _ := uuid.Parse(c.Param("id"))
		expense.ID = expenseID
	}

	if expense.ID == uuid.Nil {
		expense.ID = uuid.New()
	}
	coupleID, _ := uuid.Parse(cID)
	expense.CoupleID = coupleID
	expense, response := w.App.Cockroach.UpsertExpense(expense)
	response.JSON(c, expense)
	return nil
}

// FindExpenses find couple recurring expenses
func (w *Web) FindExpenses(c echo.Context) error {
	jwt := utils.JWTGetter(c, "couple_id")
	cID := jwt[0].(string)

	defaultParams := NewQueryParams()
	c.Bind(defaultParams)

	coupleID, _ := uuid.Parse(cID)
	expenses, response := w.App.Cockroach.FindExpenses(coupleID, defaultParams)
	response.JSON(c, expenses)
	return nil
}
