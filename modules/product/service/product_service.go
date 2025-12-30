package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
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
	productRepository repository.ProductRepository
	db                *gorm.DB
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error) {
	product := entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
		CategoryID:  req.CategoryID,
	}

	saved, err := p.productRepository.CreateProduct(ctx, p.db, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:          saved.ID.String(),
		Name:        saved.Name,
		Description: saved.Description,
		Price:       saved.Price,
		Stock:       saved.Stock,
		ImageURL:    saved.ImageURL,
		CategoryID:  saved.CategoryID,
	}, nil
}

// GetAllProducts implements ProductService.
func (p *productService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := p.productRepository.GetAllProducts(ctx, p.db)
	if err != nil {
		return nil, err
	}

	var resp []dto.ProductResponse
	for _, cat := range products {
		resp = append(resp, dto.ProductResponse{
			ID:          cat.ID.String(),
			Name:        cat.Name,
			Description: cat.Description,
			Price:       cat.Price,
			Stock:       cat.Stock,
			ImageURL:    cat.ImageURL,
			CategoryID:  cat.CategoryID,
		})
	}
	return resp, nil
}

// GetProductByID implements ProductService.
func (p *productService) GetProductByID(ctx context.Context, productId uuid.UUID) (dto.ProductResponse, error) {
	product, err := p.productRepository.GetProductByID(ctx, p.db, productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		CategoryID:  product.CategoryID,
	}, nil
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(ctx context.Context, productId uuid.UUID) error {
	return p.productRepository.DeleteProduct(ctx, p.db, productId)
}

func NewProductService(
	productsRepo repository.ProductRepository,
	db *gorm.DB,
) ProductService {
	return &productService{
		productRepository: productsRepo,
		db:                db,
	}
}
