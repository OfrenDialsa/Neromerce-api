package repository

import (
	"gorm.io/gorm"
)

type ProductsRepository interface {
}

type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) ProductsRepository {
	return &productsRepository{
		db: db,
	}
}
