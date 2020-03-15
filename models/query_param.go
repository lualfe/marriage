package models

// QueryParam models standart query params
type QueryParam struct {
	Page     int `query:"p"`
	PageSize int `query:"ps"`
}
