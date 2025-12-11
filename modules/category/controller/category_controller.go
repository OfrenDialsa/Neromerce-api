package controller

import (
	"github.com/ofrendialsa/neromerce/modules/category/service"
	"github.com/ofrendialsa/neromerce/modules/category/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	CategoryController interface {
	}

	categoryController struct {
		categoryService    service.CategoryService
		categoryValidation *validation.CategoryValidation
		db                             *gorm.DB
	}
)

func NewCategoryController(injector *do.Injector, s service.CategoryService) CategoryController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	categoryValidation := validation.NewCategoryValidation()
	return &categoryController{
		categoryService:    s,
		categoryValidation: categoryValidation,
		db:                             db,
	}
}
