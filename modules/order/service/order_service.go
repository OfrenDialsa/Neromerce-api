package service

import (
	"github.com/ofrendialsa/neromerce/modules/order/repository"
	"gorm.io/gorm"
)

type OrderService interface {
}

type orderService struct {
	orderRepository repository.OrderRepository
	db                            *gorm.DB
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	db *gorm.DB,
) OrderService {
	return &orderService{
		orderRepository: orderRepo,
		db:                            db,
	}
}
