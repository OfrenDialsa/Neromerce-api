package validation

import (
	"github.com/go-playground/validator/v10"
)

type OrderValidation struct {
	validate *validator.Validate
}

func NewOrderValidation() *OrderValidation {
	validate := validator.New()
	return &OrderValidation{
		validate: validate,
	}
}
