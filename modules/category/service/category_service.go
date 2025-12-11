package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
	"github.com/ofrendialsa/neromerce/modules/category/repository"
	"gorm.io/gorm"
)

type CategoryService interface {
	Create(ctx context.Context, req dto.CategoryCreateRequest) (dto.CategoryResponse, error)
	GetAll(ctx context.Context) ([]dto.CategoryResponse, error)
	GetCategoryById(ctx context.Context, categoryId string) (dto.CategoryResponse, error)
	Delete(ctx context.Context, categoryId string) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
	db           *gorm.DB
}

func NewCategoryService(repo repository.CategoryRepository, db *gorm.DB) CategoryService {
	return &categoryService{
		categoryRepo: repo,
		db:           db,
	}
}

func (s *categoryService) Create(ctx context.Context, req dto.CategoryCreateRequest) (dto.CategoryResponse, error) {
	category := entities.Category{
		ID:   uuid.New(),
		Name: req.Name,
	}

	saved, err := s.categoryRepo.CreateCategory(ctx, s.db, category)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:   saved.ID.String(),
		Name: saved.Name,
	}, nil
}

func (s *categoryService) GetAll(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetAllCategories(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var resp []dto.CategoryResponse
	for _, cat := range categories {
		resp = append(resp, dto.CategoryResponse{
			ID:   cat.ID.String(),
			Name: cat.Name,
		})
	}
	return resp, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, categoryId string) (dto.CategoryResponse, error) {
	category, err := s.categoryRepo.GetCategoryByID(ctx, s.db, categoryId)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:   category.ID.String(),
		Name: category.Name,
	}, nil
}

func (s *categoryService) Delete(ctx context.Context, categoryId string) error {
	return s.categoryRepo.DeleteCategory(ctx, s.db, categoryId)
}
