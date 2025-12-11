package validation

import (
	"github.com/go-playground/validator/v10"
)

type ProductsValidation struct {
	validate *validator.Validate
}

func NewProductsValidation() *ProductsValidation {
	validate := validator.New()
	return &ProductsValidation{
		validate: validate,
	}
}
