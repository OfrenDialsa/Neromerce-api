package mapper

import (
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
)

func ProductToResponse(product entities.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		CategoryID:  product.CategoryID,
	}
}
