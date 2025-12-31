package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
)

type CategoryValidation struct {
	validate *validator.Validate
}

func NewCategoryValidation() *CategoryValidation {
	validate := validator.New()

	validate.RegisterValidation("name", validateName)

	return &CategoryValidation{
		validate: validate,
	}
}

func (v *CategoryValidation) ValidateCategoryCreateRequest(req dto.CategoryCreateRequest) error {
	return v.validate.Struct(req)
}

func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return len(name) > 0 && len(name) <= 100
}
