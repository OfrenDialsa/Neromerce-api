package repository

import (
	"context"

	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error)
		GetCategoryByID(ctx context.Context, categoryId uint) (entities.Category, error)
		GetAllCategories(ctx context.Context) ([]entities.Category, error)
		DeleteCategory(ctx context.Context, categoryId uint) error
	}

	categoryRepository struct {
		db *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category entities.Category) (entities.Category, error) {
	if err := r.db.WithContext(ctx).Create(&category).Error; err != nil {
		return entities.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) GetAllCategories(ctx context.Context) ([]entities.Category, error) {

	var categories []entities.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, categoryId uint) (entities.Category, error) {
	var category entities.Category
	if err := r.db.WithContext(ctx).First(&category, categoryId).Error; err != nil {
		return entities.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, categoryId uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Category{}, categoryId).Error
}
