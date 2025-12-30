package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
	"github.com/ofrendialsa/neromerce/modules/product/repository"
	"github.com/ofrendialsa/neromerce/modules/product/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductByID(t *testing.T) {
	// 1. Setup Mock & Service
	mockRepo := new(repository.MockProductRepository)
	// Kita bisa passing nil untuk gorm.DB di service saat testing jika service tidak memanggil db.Transaction
	service := service.NewProductService(mockRepo, nil)

	ctx := context.TODO()
	productID := uuid.New()

	t.Run("Success", func(t *testing.T) {
		// 2. Ekspektasi: Jika Repo dipanggil dengan ID ini, kembalikan data produk
		expectedProduct := entities.Product{
			ID:   productID,
			Name: "Laptop Gaming",
		}
		mockRepo.On("GetProductByID", ctx, mock.Anything, productID).Return(expectedProduct, nil).Once()

		// 3. Eksekusi
		result, err := service.GetProductByID(ctx, productID)

		// 4. Asersi
		assert.NoError(t, err)
		assert.Equal(t, expectedProduct.Name, result.Name)
		assert.Equal(t, expectedProduct.ID.String(), result.ID)
		mockRepo.AssertExpectations(t) // Pastikan repo benar-benar dipanggil
	})

	t.Run("Error Not Found", func(t *testing.T) {
		// Ekspektasi: Repo mengembalikan error
		mockRepo.On("GetProductByID", ctx, mock.Anything, productID).Return(entities.Product{}, errors.New("not found")).Once()

		result, err := service.GetProductByID(ctx, productID)

		assert.Error(t, err)
		assert.Empty(t, result.Name)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllProducts(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := service.NewProductService(mockRepo, nil)
	ctx := context.TODO()

	t.Run("Success", func(t *testing.T) {
		productID := uuid.New()
		expectedEntities := []entities.Product{
			{ID: productID, Name: "Produk 1", Price: 1000},
			{ID: uuid.New(), Name: "Produk 2", Price: 2000},
		}

		mockRepo.On("GetAllProducts", ctx, mock.Anything).Return(expectedEntities, nil).Once()

		result, err := service.GetAllProducts(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Produk 1", result[0].Name)
		assert.Equal(t, productID.String(), result[0].ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Database", func(t *testing.T) {
		mockRepo.On("GetAllProducts", ctx, mock.Anything).Return([]entities.Product{}, errors.New("db error")).Once()

		result, err := service.GetAllProducts(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := service.NewProductService(mockRepo, nil)
	ctx := context.TODO()

	req := dto.ProductCreateRequest{
		Name:  "Meja Belajar",
		Price: 150000,
		Stock: 5,
	}

	t.Run("Success", func(t *testing.T) {
		productID := uuid.New()
		expectedEntity := entities.Product{
			ID:    productID,
			Name:  req.Name,
			Price: req.Price,
			Stock: req.Stock,
		}

		// Menggunakan mock.MatchedBy untuk memverifikasi data yang masuk ke repo sama dengan req
		mockRepo.On("CreateProduct", ctx, mock.Anything, mock.MatchedBy(func(p entities.Product) bool {
			return p.Name == req.Name && p.Price == req.Price
		})).Return(expectedEntity, nil).Once()

		result, err := service.CreateProduct(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, productID.String(), result.ID)
		assert.Equal(t, req.Name, result.Name)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := service.NewProductService(mockRepo, nil)
	ctx := context.TODO()
	productID := uuid.New()

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeleteProduct", ctx, mock.Anything, productID).Return(nil).Once()

		err := service.DeleteProduct(ctx, productID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepo.On("DeleteProduct", ctx, mock.Anything, productID).Return(errors.New("failed to delete")).Once()

		err := service.DeleteProduct(ctx, productID)

		assert.Error(t, err)
		assert.Equal(t, "failed to delete", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
