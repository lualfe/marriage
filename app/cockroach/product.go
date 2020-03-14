package cockroach

import (
	"strings"

	"github.com/google/uuid"
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

// CoupleIDFinder model
type CoupleIDFinder struct {
	CoupleID uuid.UUID
}

// Finder gets the list of products from a couple
func (f *CoupleIDFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("couple_id = ?", f.CoupleID).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// QueryBuild builds the cockroach query
func (f *CoupleIDFinder) QueryBuild() (string, interface{}) {
	return "couple_id = ?", f.CoupleID
}

// RoomFinder model
type RoomFinder struct {
	CoupleID uuid.UUID
	Room     string
}

// Finder gets the list of products from a room
func (f *RoomFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("couple_id = ? and room = ?", f.CoupleID, f.Room).Find(&products).Error; err != nil {
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
	CoupleID uuid.UUID
	Name     string
}

// Finder gets the list of products from a product name
func (f *NameFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("couple_id = ? and name = ?", f.CoupleID, f.Name).Find(&products).Error; err != nil {
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
	CoupleID uuid.UUID
	Brand    string
}

// Finder gets the list of products from a brand
func (f *BrandFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("couple_id = ? and brand = ?", f.CoupleID, f.Brand).Find(&products).Error; err != nil {
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
