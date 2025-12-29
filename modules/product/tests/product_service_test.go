package tests

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/ofrendialsa/neromerce/database/entities"
// 	"github.com/ofrendialsa/neromerce/modules/product/mocks"
// 	"github.com/ofrendialsa/neromerce/modules/product/service"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// func TestProductService_CreateProduct_Success(t *testing.T) {
// 	ctx := context.Background()

// 	repo := new(mocks.ProductRepositoryMock)
// 	svc := service.NewProductService(repo)

// 	product := entities.Product{
// 		Name:  "Laptop",
// 		Price: 1500,
// 	}

// 	expected := product
// 	expected.ID = uuid.New()

// 	repo.
// 		On("CreateProduct", ctx, (*gorm.DB)(nil), product).
// 		Return(expected, nil).
// 		Once()

// 	result, err := svc.CreateProduct(ctx, product)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.ID, result.ID)
// 	repo.AssertExpectations(t)
// }

// func TestProductService_CreateProduct_RepoError(t *testing.T) {
// 	ctx := context.Background()

// 	repo := new(mocks.ProductRepositoryMock)
// 	svc := service.NewProductService(repo)

// 	product := entities.Product{Name: "Laptop"}

// 	repo.
// 		On("CreateProduct", ctx, (*gorm.DB)(nil), product).
// 		Return(entities.Product{}, errors.New("db error")).
// 		Once()

// 	_, err := svc.CreateProduct(ctx, product)

// 	assert.Error(t, err)
// 	repo.AssertExpectations(t)
// }

// func TestProductService_GetAllProducts_Success(t *testing.T) {
// 	ctx := context.Background()

// 	repo := new(mocks.ProductRepositoryMock)
// 	svc := service.NewProductService(repo)

// 	expected := []entities.Product{
// 		{ID: uuid.New(), Name: "Laptop"},
// 		{ID: uuid.New(), Name: "Phone"},
// 	}

// 	repo.
// 		On("GetAllProducts", ctx, (*gorm.DB)(nil)).
// 		Return(expected, nil).
// 		Once()

// 	products, err := svc.GetAllProducts(ctx)

// 	assert.NoError(t, err)
// 	assert.Len(t, products, 2)
// 	repo.AssertExpectations(t)
// }

// func TestProductService_DeleteProduct_Success(t *testing.T) {
// 	ctx := context.Background()

// 	repo := new(mocks.ProductRepositoryMock)
// 	svc := service.NewProductService(repo)

// 	id := uuid.New()

// 	repo.
// 		On("DeleteProduct", ctx, (*gorm.DB)(nil), id).
// 		Return(nil).
// 		Once()

// 	err := svc.DeleteProduct(ctx, id)

// 	assert.NoError(t, err)
// 	repo.AssertExpectations(t)
// }
