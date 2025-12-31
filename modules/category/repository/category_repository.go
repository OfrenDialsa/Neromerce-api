package repository

import (
	"context"
	"strings"

	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
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

	err := r.db.WithContext(ctx).Create(&category).Error
	if err != nil {

		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE") {
			return entities.Category{}, dto.ErrCategoryNameExist
		}

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
	result := r.db.WithContext(ctx).
		Delete(&entities.Category{}, categoryId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
