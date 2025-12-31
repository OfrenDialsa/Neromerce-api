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
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, tx *gorm.DB, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, tx, product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllProducts(ctx context.Context, tx *gorm.DB) ([]entities.Product, error) {
	args := m.Called(ctx, tx)
	// Kita perlu handle return slice agar tidak panic jika nil
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, tx *gorm.DB, productId uuid.UUID) error {
	args := m.Called(ctx, tx, productId)
	return args.Error(0)
}
