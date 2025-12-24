package controller

import (
	"github.com/ofrendialsa/neromerce/modules/product/service"
	"github.com/ofrendialsa/neromerce/modules/product/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	ProductController interface {
	}

	productController struct {
		productService    service.ProductService
		productValidation *validation.ProductValidation
		db                *gorm.DB
	}
)

func NewProductsController(injector *do.Injector, s service.ProductService) ProductController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	productValidation := validation.NewProductValidation()
	return &productController{
		productService:    s,
		productValidation: productValidation,
		db:                db,
	}
}
