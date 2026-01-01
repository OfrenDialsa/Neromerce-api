package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
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

func (v *ProductValidation) ValidateCreateRequest(req dto.ProductCreateRequest) error {
	return v.validate.Struct(req)
}
