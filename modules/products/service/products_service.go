package service

import (
	"github.com/ofrendialsa/neromerce/modules/products/repository"
	"gorm.io/gorm"
)

type ProductsService interface {
}

type productsService struct {
	productsRepository repository.ProductsRepository
	db                            *gorm.DB
}

func NewProductsService(
	productsRepo repository.ProductsRepository,
	db *gorm.DB,
) ProductsService {
	return &productsService{
		productsRepository: productsRepo,
		db:                            db,
	}
}
