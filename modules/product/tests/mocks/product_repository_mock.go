package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) CreateProduct(
	ctx context.Context,
	tx *gorm.DB,
	product entities.Product,
) (entities.Product, error) {

	args := m.Called(ctx, tx, product)

	var result entities.Product
	if args.Get(0) != nil {
		result = args.Get(0).(entities.Product)
	}

	return result, args.Error(1)
}

func (m *ProductRepositoryMock) GetProductByID(
	ctx context.Context,
	tx *gorm.DB,
	productId uuid.UUID,
) (entities.Product, error) {

	args := m.Called(ctx, tx, productId)

	var result entities.Product
	if args.Get(0) != nil {
		result = args.Get(0).(entities.Product)
	}

	return result, args.Error(1)
}

func (m *ProductRepositoryMock) GetAllProducts(
	ctx context.Context,
	tx *gorm.DB,
) ([]entities.Product, error) {

	args := m.Called(ctx, tx)

	var result []entities.Product
	if args.Get(0) != nil {
		result = args.Get(0).([]entities.Product)
	}

	return result, args.Error(1)
}

func (m *ProductRepositoryMock) DeleteProduct(
	ctx context.Context,
	tx *gorm.DB,
	productId uuid.UUID,
) error {

	args := m.Called(ctx, tx, productId)
	return args.Error(0)
}
