package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
	"github.com/ofrendialsa/neromerce/modules/category/service"
	"github.com/ofrendialsa/neromerce/modules/category/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/ofrendialsa/neromerce/pkg/utils"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	CategoryController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetCategoryByID(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	categoryController struct {
		categoryService    service.CategoryService
		categoryValidation *validation.CategoryValidation
		db                 *gorm.DB
	}
)

func NewCategoryController(injector *do.Injector, s service.CategoryService) CategoryController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	categoryValidation := validation.NewCategoryValidation()
	return &categoryController{
		categoryService:    s,
		categoryValidation: categoryValidation,
		db:                 db,
	}
}

// Create implements CategoryController.
func (c *categoryController) Create(ctx *gin.Context) {
	var req dto.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Validasi
	if err := c.categoryValidation.ValidateCategoryCreateRequest(req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CATEGORY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	category, err := c.categoryService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CATEGORY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_CATEGORY, category)
	ctx.JSON(http.StatusOK, res)
}

// GetAll implements CategoryController.
func (c *categoryController) GetAll(ctx *gin.Context) {
	categories, err := c.categoryService.GetAll(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_CATEGORY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_CATEGORY, categories)
	ctx.JSON(http.StatusOK, res)
}

// GetCategoryByID implements CategoryController.
func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	category, err := c.categoryService.GetCategoryById(ctx.Request.Context(), categoryId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_CATEGORY, err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_CATEGORY, category)
	ctx.JSON(http.StatusOK, res)
}

// Delete implements CategoryController.
func (c *categoryController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("id")

	if err := c.categoryService.Delete(ctx.Request.Context(), categoryId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_CATEGORY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_CATEGORY, nil)
	ctx.JSON(http.StatusOK, res)
}
