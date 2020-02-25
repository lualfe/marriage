package cockroach

import "github.com/lualfe/casamento/models"

// FindProducts gets all the products
func (a *DB) FindProducts() ([]*models.Product, error) {
	d := []*models.Product{}
	if err := a.Instance.Find(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}
