package web

import (
	"github.com/lualfe/casamento/models"
)

// NewQueryParams initializes a Query params default struct
func NewQueryParams() *models.QueryParam {
	qp := &models.QueryParam{1, 10}
	return qp
}
