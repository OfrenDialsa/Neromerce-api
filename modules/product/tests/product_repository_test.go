package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/config"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/product/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db := config.SetUpTestDatabaseConnection()

	if err := db.Migrator().DropTable(&entities.Product{}, &entities.Category{}); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}

	if err := db.AutoMigrate(&entities.Category{}, &entities.Product{}); err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestProductRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB(t)
	defer config.CloseDatabaseConnection(db)

	repo := repository.NewProductRepository(db)

	category := entities.Category{
		Name: "Electronics",
	}
	if err := db.Create(&category).Error; err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	tests := []struct {
		name       string
		product    entities.Product
		shouldFail bool
	}{
		{
			name: "Valid product",
			product: entities.Product{
				Name:        "Laptop",
				Description: "Gaming Laptop",
				Price:       1500,
				Stock:       5,
				CategoryID:  category.ID,
			},
			shouldFail: false,
		},
		{
			name: "Empty name (allowed in repository)",
			product: entities.Product{
				Name:        "",
				Description: "No name",
				Price:       100,
				Stock:       1,
				CategoryID:  category.ID,
			},
			shouldFail: false,
		},
		{
			name: "Negative price (allowed in repository)",
			product: entities.Product{
				Name:        "Cheap Product",
				Description: "Negative price",
				Price:       -50,
				Stock:       1,
				CategoryID:  category.ID,
			},
			shouldFail: false,
		},
		{
			name: "Stock default",
			product: entities.Product{
				Name:        "Default Stock",
				Description: "Stock should default 0",
				Price:       200,
				CategoryID:  category.ID,
			},
			shouldFail: false,
		},
		{
			name: "Invalid category (FK constraint)",
			product: entities.Product{
				Name:        "Invalid Category",
				Description: "CategoryID not exist",
				Price:       100,
				Stock:       1,
				CategoryID:  9999,
			},
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			created, err := repo.CreateProduct(ctx, nil, tt.product)

			if tt.shouldFail {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEqual(t, uuid.Nil, created.ID)
		})
	}

	products, err := repo.GetAllProducts(ctx, nil)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(products), 1)

	for _, p := range products {
		err := repo.DeleteProduct(ctx, nil, p.ID)
		assert.NoError(t, err)
	}

	productsAfterDelete, err := repo.GetAllProducts(ctx, nil)
	assert.NoError(t, err)
	assert.Len(t, productsAfterDelete, 0)

	t.Run("Update product", func(t *testing.T) {
		original := entities.Product{
			Name:        "Old Name",
			Description: "Old Description",
			Price:       100,
			Stock:       10,
			CategoryID:  category.ID,
		}

		created, err := repo.CreateProduct(ctx, nil, original)
		assert.NoError(t, err)

		updateData := map[string]interface{}{
			"name":        "New Name",
			"description": "New Description",
			"price":       200.0,
			"stock":       5,
			"category_id": category.ID,
		}

		updated, err := repo.UpdateProduct(ctx, nil, created.ID, updateData)
		assert.NoError(t, err)

		assert.Equal(t, created.ID, updated.ID)
		assert.Equal(t, "New Name", updated.Name)
		assert.Equal(t, "New Description", updated.Description)
		assert.Equal(t, 200.0, updated.Price)
		assert.Equal(t, 5, updated.Stock)
	})

	t.Run("Update product not found", func(t *testing.T) {
		_, err := repo.UpdateProduct(
			ctx,
			nil,
			uuid.New(),
			map[string]interface{}{
				"name": "Does not exist",
			},
		)

		assert.Error(t, err)
	})

}
