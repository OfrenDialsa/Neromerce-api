package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductByID(ctx context.Context, tx *gorm.DB, productId uuid.UUID) (entities.Product, error) {
	args := m.Called(ctx, tx, productId)
	if args.Get(0) == nil {
		return entities.Product{}, args.Error(1)
	}

	product, ok := args.Get(0).(entities.Product)
	if !ok {
		return entities.Product{}, args.Error(1)
	}

	return product, args.Error(1)
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, tx, product)
	if args.Get(0) == nil {
		return entities.Product{}, args.Error(1)
	}

	res, ok := args.Get(0).(entities.Product)
	if !ok {
		return entities.Product{}, args.Error(1)
	}

	return res, args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID, updates map[string]interface{}) (entities.Product, error) {
	args := m.Called(ctx, tx, productId, updates)
	if args.Get(0) == nil {
		return entities.Product{}, args.Error(1)
	}

	res, ok := args.Get(0).(entities.Product)
	if !ok {
		return entities.Product{}, args.Error(1)
	}

	return res, args.Error(1)
}

func (m *MockProductRepository) GetAllProducts(ctx context.Context, tx *gorm.DB) ([]entities.Product, error) {
	args := m.Called(ctx, tx)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	products, ok := args.Get(0).([]entities.Product)
	if !ok {
		return nil, args.Error(1)
	}

	return products, args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID) error {
	args := m.Called(ctx, tx, productId)
	return args.Error(0)
}
