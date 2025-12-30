package service

import (
	"context"
	"strconv"

	"github.com/ofrendialsa/neromerce/database/entities"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
	"github.com/ofrendialsa/neromerce/modules/category/repository"
)

type CategoryService interface {
	Create(ctx context.Context, req dto.CategoryCreateRequest) (dto.CategoryResponse, error)
	GetAll(ctx context.Context) ([]dto.CategoryResponse, error)
	GetCategoryById(ctx context.Context, categoryId uint) (dto.CategoryResponse, error)
	Delete(ctx context.Context, categoryId uint) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: repo,
	}
}

func (s *categoryService) Create(ctx context.Context, req dto.CategoryCreateRequest) (dto.CategoryResponse, error) {
	category := entities.Category{
		Name: req.Name,
	}

	saved, err := s.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:   strconv.Itoa(int(saved.ID)),
		Name: saved.Name,
	}, nil
}

func (s *categoryService) GetAll(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	var resp []dto.CategoryResponse
	for _, cat := range categories {
		resp = append(resp, dto.CategoryResponse{
			ID:   strconv.Itoa(int(cat.ID)),
			Name: cat.Name,
		})
	}
	return resp, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, categoryId uint) (dto.CategoryResponse, error) {
	category, err := s.categoryRepo.GetCategoryByID(ctx, categoryId)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.CategoryResponse{
		ID:   strconv.Itoa(int(category.ID)),
		Name: category.Name,
	}, nil
}

func (s *categoryService) Delete(ctx context.Context, categoryId uint) error {
	return s.categoryRepo.DeleteCategory(ctx, categoryId)
}
