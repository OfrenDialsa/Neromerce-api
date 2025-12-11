package controller

import (
	"github.com/ofrendialsa/neromerce/modules/products/service"
	"github.com/ofrendialsa/neromerce/modules/products/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	ProductsController interface {
	}

	productsController struct {
		productsService    service.ProductsService
		productsValidation *validation.ProductsValidation
		db                             *gorm.DB
	}
)

func NewProductsController(injector *do.Injector, s service.ProductsService) ProductsController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	productsValidation := validation.NewProductsValidation()
	return &productsController{
		productsService:    s,
		productsValidation: productsValidation,
		db:                             db,
	}
}
