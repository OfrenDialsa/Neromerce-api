package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
	"github.com/ofrendialsa/neromerce/modules/product/mapper"
	"github.com/ofrendialsa/neromerce/modules/product/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error)
	GetAll(ctx context.Context) ([]dto.ProductResponse, error)
	GetByID(ctx context.Context, productId uuid.UUID) (dto.ProductResponse, error)
	Update(ctx context.Context, req dto.ProductUpdateRequest, productId uuid.UUID) (dto.ProductResponse, error)
	Delete(ctx context.Context, productId uuid.UUID) error
}

type productService struct {
	productRepository repository.ProductRepository
	db                *gorm.DB
}

// CreateProduct implements ProductService.
func (p *productService) Create(ctx context.Context, req dto.ProductCreateRequest) (dto.ProductResponse, error) {
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

	return mapper.ProductToResponse(saved), nil
}

// UpdateProduct implements ProductService.
func (p *productService) Update(ctx context.Context, req dto.ProductUpdateRequest, productId uuid.UUID) (dto.ProductResponse, error) {

	updates := map[string]interface{}{}

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.Stock != nil {
		updates["stock"] = *req.Stock
	}
	if req.ImageURL != nil {
		updates["image_url"] = *req.ImageURL
	}
	if req.CategoryID != nil {
		updates["category_id"] = *req.CategoryID
	}

	if len(updates) == 0 {
		return dto.ProductResponse{}, dto.ErrUpdateProduct
	}

	updated, err := p.productRepository.UpdateProduct(ctx, p.db, productId, updates)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return mapper.ProductToResponse(updated), nil
}

// GetAllProducts implements ProductService.
func (p *productService) GetAll(ctx context.Context) ([]dto.ProductResponse, error) {
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
func (p *productService) GetByID(ctx context.Context, productId uuid.UUID) (dto.ProductResponse, error) {
	product, err := p.productRepository.GetProductByID(ctx, p.db, productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return mapper.ProductToResponse(product), nil
}

// DeleteProduct implements ProductService.
func (p *productService) Delete(ctx context.Context, productId uuid.UUID) error {
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
