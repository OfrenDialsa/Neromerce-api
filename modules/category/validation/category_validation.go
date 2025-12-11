package validation

import (
	"github.com/go-playground/validator/v10"
)

type CategoryValidation struct {
	validate *validator.Validate
}

func NewCategoryValidation() *CategoryValidation {
	validate := validator.New()
	return &CategoryValidation{
		validate: validate,
	}
}
