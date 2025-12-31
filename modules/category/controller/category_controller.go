package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ofrendialsa/neromerce/modules/category/dto"
	"github.com/ofrendialsa/neromerce/modules/category/service"
	"github.com/ofrendialsa/neromerce/modules/category/validation"
	"github.com/ofrendialsa/neromerce/pkg/utils"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type categoryController struct {
	categoryService    service.CategoryService
	categoryValidation *validation.CategoryValidation
}

func NewCategoryController(s service.CategoryService) CategoryController {
	return &categoryController{
		categoryService:    s,
		categoryValidation: validation.NewCategoryValidation(),
	}
}

func (c *categoryController) Create(ctx *gin.Context) {
	var req dto.CategoryCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_DATA_FROM_BODY,
			"request format invalid",
			nil,
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.categoryValidation.ValidateCategoryCreateRequest(req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			e := ve[0]
			var msg string
			switch e.Tag() {
			case "required":
				msg = "category name is required"
			case "name":
				msg = "category name must be at most 100 characters"
			}
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_CATEGORY, msg, nil)
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
	}

	category, err := c.categoryService.Create(ctx.Request.Context(), req)
	if err != nil {

		switch {
		case errors.Is(err, dto.ErrCategoryNameExist):
			res := utils.BuildResponseFailed(
				dto.MESSAGE_FAILED_CREATE_CATEGORY,
				"category already exists",
				nil,
			)
			ctx.JSON(http.StatusConflict, res)
			return

		default:
			res := utils.BuildResponseFailed(
				dto.MESSAGE_FAILED_CREATE_CATEGORY,
				"internal server error",
				nil,
			)
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_CREATE_CATEGORY,
		category,
	)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) GetAll(ctx *gin.Context) {
	categories, err := c.categoryService.GetAll(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_LIST_CATEGORY,
			err.Error(),
			nil,
		)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_GET_LIST_CATEGORY,
		categories,
	)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	categoryId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_CATEGORY,
			"invalid category id",
			nil,
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	category, err := c.categoryService.GetCategoryById(
		ctx.Request.Context(),
		uint(categoryId),
	)
	if err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_CATEGORY,
			err.Error(),
			nil,
		)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_GET_CATEGORY,
		category,
	)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")

	categoryId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_DELETE_CATEGORY,
			"invalid category id",
			nil,
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.categoryService.Delete(
		ctx.Request.Context(),
		uint(categoryId),
	); err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_DELETE_CATEGORY,
			"category id not found",
			nil,
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_DELETE_CATEGORY,
		nil,
	)
	ctx.JSON(http.StatusOK, res)
}
