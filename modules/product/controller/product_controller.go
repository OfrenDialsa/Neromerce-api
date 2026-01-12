package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ofrendialsa/neromerce/modules/product/dto"
	"github.com/ofrendialsa/neromerce/modules/product/service"
	"github.com/ofrendialsa/neromerce/modules/product/validation"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/ofrendialsa/neromerce/pkg/utils"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	ProductController interface {
		Create(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	productController struct {
		productService    service.ProductService
		productValidation *validation.ProductValidation
		db                *gorm.DB
	}
)

func (p *productController) Create(ctx *gin.Context) {
	var req dto.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_DATA_FROM_BODY,
			"request format invalid",
			nil,
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := p.productValidation.ValidateCreateRequest(req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			e := ve[0]
			var msg string
			switch e.Tag() {
			case "required":
				msg = "product name is required"
			case "name":
				msg = "product name must be at most 100 characters"
			}
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PRODUCT, msg, nil)
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
	}

	product, err := p.productService.Create(ctx.Request.Context(), req)
	if err != nil {

		switch {
		case errors.Is(err, dto.ErrProductNameExist):
			res := utils.BuildResponseFailed(
				dto.MESSAGE_FAILED_CREATE_PRODUCT,
				"product already exists",
				nil,
			)
			ctx.JSON(http.StatusConflict, res)
			return

		default:
			res := utils.BuildResponseFailed(
				dto.MESSAGE_FAILED_CREATE_PRODUCT,
				"internal server error",
				nil,
			)
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_CREATE_PRODUCT,
		product,
	)
	ctx.JSON(http.StatusOK, res)

}

// Delete implements ProductController.
func (p *productController) Delete(ctx *gin.Context) {
	panic("unimplemented")
}

// GetAll implements ProductController.
func (p *productController) GetAll(ctx *gin.Context) {
	products, err := p.productService.GetAll(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(
			dto.MESSAGE_FAILED_GET_LIST_PRODUCT,
			err.Error(),
			nil,
		)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(
		dto.MESSAGE_SUCCESS_GET_LIST_PRODUCT,
		products,
	)
	ctx.JSON(http.StatusOK, res)
}

// GetById implements ProductController.
func (p *productController) GetById(ctx *gin.Context) {
	panic("unimplemented")
}

// Update implements ProductController.
func (p *productController) Update(ctx *gin.Context) {
	panic("unimplemented")
}

func NewProductController(injector *do.Injector, s service.ProductService) ProductController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	productValidation := validation.NewProductValidation()
	return &productController{
		productService:    s,
		productValidation: productValidation,
		db:                db,
	}
}
