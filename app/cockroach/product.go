package cockroach

import (
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
	"github.com/lualfe/casamento/utils"
)

// ProductFinder is responsible for finding products
type ProductFinder interface {
	Finder(a *DB) ([]*models.Product, responsewriter.Response)
	CompareField(now, before []*models.Product) []*models.Product
}

// IDFinder model
type IDFinder struct {
	ID string
}

// Finder gets the list of products from a user
func (f *IDFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ?", f.ID).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// CompareField compares a field on two different products
func (f *IDFinder) CompareField(now, before []*models.Product) []*models.Product {
	products := []*models.Product{}
	for _, n := range now {
		for _, b := range before {
			if n.ID == b.ID {
				products = append(products, b)
			}
		}
	}
	return products
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

// CompareField compares a field on two different products
func (f *RoomFinder) CompareField(now, before []*models.Product) []*models.Product {
	products := []*models.Product{}
	for _, n := range now {
		for _, b := range before {
			if n.Room == b.Room {
				products = append(products, b)
			}
		}
	}
	return products
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

// CompareField compares a field on two different products
func (f *NameFinder) CompareField(now, before []*models.Product) []*models.Product {
	products := []*models.Product{}
	for _, n := range now {
		for _, b := range before {
			if n.Name == b.Name {
				products = append(products, b)
			}
		}
	}
	return products
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

// CompareField compares a field on two different products
func (f *BrandFinder) CompareField(now, before []*models.Product) []*models.Product {
	products := []*models.Product{}
	for _, n := range now {
		for _, b := range before {
			if n.Brand == b.Brand {
				products = append(products, b)
			}
		}
	}
	return products
}

// MultipleFinder model
type MultipleFinder struct {
	Multiple []ProductFinder
}

// Finder gets the list of products from a multiple fields
func (f *MultipleFinder) Finder(a *DB) ([]*models.Product, responsewriter.Response) {
	beforeProducts := []*models.Product{}
	products := []*models.Product{}
	for _, v := range f.Multiple {
		nowProducts, _ := v.Finder(a)
		products = v.CompareField(nowProducts, beforeProducts)
		beforeProducts = products
		if len(products) == 0 {
			beforeProducts = nowProducts
		}
	}
	noDuplicates := utils.Unique(products)
	return noDuplicates, responsewriter.Success()
}

// CompareField compares a field on two different products
func (f *MultipleFinder) CompareField(now, before []*models.Product) []*models.Product {
	products := []*models.Product{}
	for _, n := range now {
		for _, b := range before {
			if n == b {
				products = append(products, b)
			}
		}
	}
	return products
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
