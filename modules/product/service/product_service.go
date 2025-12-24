package service

import (
	"github.com/ofrendialsa/neromerce/modules/product/repository"
	"gorm.io/gorm"
)

type ProductService interface {
}

type productsService struct {
	productsRepository repository.ProductRepository
	db                 *gorm.DB
}

func NewProductsService(
	productsRepo repository.ProductRepository,
	db *gorm.DB,
) ProductService {
	return &productsService{
		productsRepository: productsRepo,
		db:                 db,
	}
}
