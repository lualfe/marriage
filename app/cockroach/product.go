package cockroach

import (
	"github.com/lualfe/casamento/app/responsewriter"
	"github.com/lualfe/casamento/models"
)

// FindProducts gets all the products
func (a *DB) FindProducts(id string) ([]*models.Product, responsewriter.Response) {
	products := []*models.Product{}
	if err := a.Instance.Where("user_id = ?", id).Find(&products).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return products, responsewriter.Success()
}

// UpsertProduct upserts the product
func (a *DB) UpsertProduct(product *models.Product) (*models.Product, responsewriter.Response) {
	if err := a.Instance.Save(&product).Error; err != nil {
		return nil, responsewriter.UnexpectedError(err)
	}
	return product, responsewriter.Success()
}
