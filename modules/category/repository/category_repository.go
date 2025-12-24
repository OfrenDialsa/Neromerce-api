package repository

import (
	"context"

	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		CreateCategory(ctx context.Context, tx *gorm.DB, category entities.Category) (entities.Category, error)
		GetCategoryByID(ctx context.Context, tx *gorm.DB, categoryId uint) (entities.Category, error)
		GetAllCategories(ctx context.Context, tx *gorm.DB) ([]entities.Category, error)
		DeleteCategory(ctx context.Context, tx *gorm.DB, categoryId uint) error
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

func (c *categoryRepository) CreateCategory(ctx context.Context, tx *gorm.DB, category entities.Category) (entities.Category, error) {
	if tx == nil {
		tx = c.db
	}

	if err := tx.WithContext(ctx).Create(&category).Error; err != nil {
		return entities.Category{}, err
	}

	return category, nil
}

func (c *categoryRepository) GetAllCategories(ctx context.Context, tx *gorm.DB) ([]entities.Category, error) {
	if tx == nil {
		tx = c.db
	}

	var categories []entities.Category
	if err := tx.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryRepository) GetCategoryByID(ctx context.Context, tx *gorm.DB, categoryId uint) (entities.Category, error) {
	if tx == nil {
		tx = c.db
	}

	var category entities.Category
	if err := tx.WithContext(ctx).First(&category, categoryId).Error; err != nil {
		return entities.Category{}, err
	}

	return category, nil
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, tx *gorm.DB, categoryId uint) error {
	if tx == nil {
		tx = c.db
	}

	if err := tx.WithContext(ctx).Delete(&entities.Category{}, categoryId).Error; err != nil {
		return err
	}

	return nil
}
