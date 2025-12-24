package validation

import (
	"github.com/go-playground/validator/v10"
)

type ProductValidation struct {
	validate *validator.Validate
}

func NewProductValidation() *ProductValidation {
	validate := validator.New()
	return &ProductValidation{
		validate: validate,
	}
}
