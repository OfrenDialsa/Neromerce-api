package controller

import (
	"github.com/ofrendialsa/neromerce/modules/order/service"
	"github.com/ofrendialsa/neromerce/modules/order/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	OrderController interface {
	}

	orderController struct {
		orderService    service.OrderService
		orderValidation *validation.OrderValidation
		db                             *gorm.DB
	}
)

func NewOrderController(injector *do.Injector, s service.OrderService) OrderController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	orderValidation := validation.NewOrderValidation()
	return &orderController{
		orderService:    s,
		orderValidation: orderValidation,
		db:                             db,
	}
}
