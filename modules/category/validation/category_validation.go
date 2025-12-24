package validation

import (
	"fmt"

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
	err := v.validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validationErrors {
				switch ve.Tag() {
				case "required":
					return fmt.Errorf("%s wajib diisi", ve.Field())
				case "max":
					return fmt.Errorf("%s maksimal %s karakter", ve.Field(), ve.Param())
				case "min":
					return fmt.Errorf("%s minimal %s karakter", ve.Field(), ve.Param())
				case "name":
					return fmt.Errorf("%s tidak valid", ve.Field())
				}
			}
		}
		return err
	}
	return nil
}

func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	return len(name) > 0 && len(name) <= 100
}
