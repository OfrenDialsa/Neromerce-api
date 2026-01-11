package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		CreateProduct(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error)
		GetProductByID(ctx context.Context, tx *gorm.DB, productId uuid.UUID) (entities.Product, error)
		GetAllProducts(ctx context.Context, tx *gorm.DB) ([]entities.Product, error)
		UpdateProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID, updates map[string]interface{}) (entities.Product, error)
		DeleteProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID) error
	}

	productRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

// CreateProduct implements ProductRepository
func (p *productRepository) CreateProduct(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error) {
	if tx == nil {
		tx = p.db
	}

	if product.ID == uuid.Nil {
		product.ID = uuid.New()
	}

	if err := tx.WithContext(ctx).Create(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

// UpdateProduct implements ProductRepository.
func (p *productRepository) UpdateProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID, updates map[string]interface{}) (entities.Product, error) {

	if tx == nil {
		tx = p.db
	}

	if len(updates) == 0 {
		return entities.Product{}, gorm.ErrInvalidData
	}

	if err := tx.WithContext(ctx).
		Model(&entities.Product{}).
		Where("id = ?", productId).
		Updates(updates).Error; err != nil {
		return entities.Product{}, err
	}

	var product entities.Product
	if err := tx.WithContext(ctx).
		First(&product, "id = ?", productId).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

// GetAllProducts implements ProductRepository
func (p *productRepository) GetAllProducts(ctx context.Context, tx *gorm.DB) ([]entities.Product, error) {
	if tx == nil {
		tx = p.db
	}

	var products []entities.Product
	if err := tx.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID implements ProductRepository
func (p *productRepository) GetProductByID(ctx context.Context, tx *gorm.DB, productId uuid.UUID) (entities.Product, error) {
	if tx == nil {
		tx = p.db
	}

	var product entities.Product
	if err := tx.WithContext(ctx).Where("id = ?", productId).First(&product).Error; err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

// DeleteProduct implements ProductRepository
func (p *productRepository) DeleteProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID) error {
	if tx == nil {
		tx = p.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", productId).Delete(&entities.Product{}).Error; err != nil {
		return err
	}

	return nil
}
