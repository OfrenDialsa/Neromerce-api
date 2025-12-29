package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
	"github.com/ofrendialsa/neromerce/modules/product/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error)
	GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error)
	GetProductByID(ctx context.Context, productId uuid.UUID) (dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, productId uuid.UUID) error
}

type productService struct {
	productsRepository repository.ProductRepository
	db                 *gorm.DB
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error) {
	panic("unimplemented")
}

// GetAllProducts implements ProductService.
func (p *productService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	panic("unimplemented")
}

// GetProductByID implements ProductService.
func (p *productService) GetProductByID(ctx context.Context, productId uuid.UUID) (dto.ProductResponse, error) {
	panic("unimplemented")
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	panic("unimplemented")
}

func NewProductService(
	productsRepo repository.ProductRepository,
	db *gorm.DB,
) ProductService {
	return &productService{
		productsRepository: productsRepo,
		db:                 db,
	}
}
