package cockroach

import (
	"github.com/google/uuid"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
)

// UpsertMonthlyIncome inserts an income into the database. If already exists, update
func (a *DB) UpsertMonthlyIncome(income *models.MonthlyIncome) (*models.MonthlyIncome, responsewriter.Response) {
	if err := a.Instance.Save(&income).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return income, responsewriter.Success()
}

// FindMonthlyIncome finds a couple income
func (a *DB) FindMonthlyIncome(coupleID uuid.UUID) (*models.MonthlyIncome, responsewriter.Response) {
	income := &models.MonthlyIncome{}
	if err := a.Instance.Where("couple_id = ?", coupleID).First(&income).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return income, responsewriter.Success()
}
