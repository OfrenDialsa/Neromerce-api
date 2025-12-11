package service

import (
	"github.com/ofrendialsa/neromerce/modules/category/repository"
	"gorm.io/gorm"
)

type CategoryService interface {
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
	db                            *gorm.DB
}

func NewCategoryService(
	categoryRepo repository.CategoryRepository,
	db *gorm.DB,
) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
		db:                            db,
	}
}
