package utils

import (
	"github.com/lualfe/casamento/models"
)

// Unique is a helper function that eliminates duplicate products from an products slice
func Unique(productsSlice []*models.Product) []*models.Product {
	keys := make(map[string]bool)
	list := []*models.Product{}
	for _, product := range productsSlice {
		if _, value := keys[product.ID.String()]; !value {
			keys[product.ID.String()] = true
			list = append(list, product)
		}
	}
	return list
}
