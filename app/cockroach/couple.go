package cockroach

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// CreateCouple creates a new couple
func (a *DB) CreateCouple() (uuid.UUID, responsewriter.Response) {
	couple := &models.Couple{
		ID:    uuid.New(),
		Token: utils.RandStringRunes(10),
	}
	if err := a.Instance.Save(&couple).Error; err != nil {
		return uuid.Nil, responsewriter.UnexpectedError(err)
	}
	return couple.ID, responsewriter.Success()
}

// CheckCoupleToken checks if a token is valid
func (a *DB) CheckCoupleToken(token string, coupleID uuid.UUID) (bool, responsewriter.Response) {
	couple := &models.Couple{}
	if err := a.Instance.Where("id = ?", coupleID).Find(&couple).Error; err != nil {
		return false, responsewriter.UnexpectedError(err)
	}
	if couple.Token != token {
		return false, responsewriter.Error("token incorrect", http.StatusUnauthorized)
	}
	return true, responsewriter.Success()
}

// DeleteCouple deletes an couple by id
func (a *DB) DeleteCouple(coupleID uuid.UUID) responsewriter.Response {
	if err := a.Instance.Unscoped().Delete(&models.Couple{}, "id = ?", coupleID).Error; err != nil {
		return responsewriter.UnexpectedError(err)
	}
	return responsewriter.Success()
}
