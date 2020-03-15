package cockroach

import (
	"github.com/google/uuid"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
)

// UpsertExpense inserts an expense in database. If already existis, updates
func (a *DB) UpsertExpense(expense *models.RecurringExpense) (*models.RecurringExpense, responsewriter.Response) {
	if err := a.Instance.Save(&expense).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return expense, responsewriter.Success()
}

// FindExpenses finds couple recurring expenses
func (a *DB) FindExpenses(coupleID uuid.UUID, params *models.QueryParam) ([]*models.RecurringExpense, responsewriter.Response) {
	expenses := []*models.RecurringExpense{}
	if err := a.Instance.Where("couple_id = ?", coupleID).
		Offset((params.Page - 1) * params.PageSize).
		Limit(params.PageSize).
		Find(&expenses).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return expenses, responsewriter.Success()
}
