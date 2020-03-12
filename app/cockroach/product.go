package cockroach

import (
	"strings"

	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
)

// ProductGather gathers all products related interfaces
type ProductGather interface {
	ProductFinder
	ProductsRepository
}

// ProductFinder is responsible for finding products
type ProductFinder interface {
	Finder(a *DB) ([]*models.Product, responsewriter.Response)
}

// ProductsRepository is responsible for comparing products
type ProductsRepository interface {
	QueryBuild() (string, interface{})
}

// UserIDFinder model
type UserIDFinder struct {
	UserID string
}

// Finder gets the list of products from a user
func (f *UserIDFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ?", f.UserID).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// QueryBuild builds the cockroach query
func (f *UserIDFinder) QueryBuild() (string, interface{}) {
	return "user_id = ?", f.UserID
}

// RoomFinder model
type RoomFinder struct {
	UserID string
	Room   string
}

// Finder gets the list of products from a room
func (f *RoomFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ? and room = ?", f.UserID, f.Room).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// QueryBuild builds the cockroach query
func (f *RoomFinder) QueryBuild() (string, interface{}) {
	return "LOWER(room) = LOWER(?)", f.Room
}

// NameFinder model
type NameFinder struct {
	UserID string
	Name   string
}

// Finder gets the list of products from a product name
func (f *NameFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ? and name = ?", f.UserID, f.Name).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// QueryBuild builds the cockroach query
func (f *NameFinder) QueryBuild() (string, interface{}) {
	return "LOWER(name) = LOWER(?)", f.Name
}

// BrandFinder model
type BrandFinder struct {
	UserID string
	Brand  string
}

// Finder gets the list of products from a brand
func (f *BrandFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ? and brand = ?", f.UserID, f.Brand).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// QueryBuild builds the cockroach query
func (f *BrandFinder) QueryBuild() (string, interface{}) {
	return "LOWER(brand) = LOWER(?)", f.Brand
}

// MultipleFinder model
type MultipleFinder struct {
	Multiple []ProductGather
}

// Finder gets the list of products from a multiple fields
func (f *MultipleFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	var gatherQueries []string
	var gatherParams []interface{}
	for _, v := range f.Multiple {
		query, param := v.QueryBuild()
		gatherQueries = append(gatherQueries, query)
		gatherParams = append(gatherParams, param)
	}
	finalQuery := strings.Join(gatherQueries, " AND ")
	if err := a.Instance.Where(finalQuery, gatherParams...).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// FindProducts gets all the products
func (a *DB) FindProducts(finder ProductFinder) ([]*models.Product, responsewriter.Response) {
	products, response := finder.Finder(a)
	return products, response
}

// UpsertProduct upserts the product
func (a *DB) UpsertProduct(product *models.Product) (*models.Product, responsewriter.Response) {
	if err := a.Instance.Save(&product).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return product, responsewriter.Success()
}
