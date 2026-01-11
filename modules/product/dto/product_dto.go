package dto

import "errors"

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA          = "success get data"
)

const (
	MESSAGE_FAILED_CREATE_PRODUCT   = "failed create product"
	MESSAGE_FAILED_GET_PRODUCT      = "failed get product"
	MESSAGE_FAILED_GET_LIST_PRODUCT = "failed get list product"
	MESSAGE_FAILED_DELETE_PRODUCT   = "failed delete product"
	MESSAGE_FAILED_PROSES_REQUEST   = "failed proses request"

	// Success
	MESSAGE_SUCCESS_CREATE_PRODUCT   = "success create product"
	MESSAGE_SUCCESS_GET_PRODUCT      = "success get product"
	MESSAGE_SUCCESS_GET_LIST_PRODUCT = "success get list product"
	MESSAGE_SUCCESS_DELETE_PRODUCT   = "success delete product"
)

var (
	ErrCreateProduct    = errors.New("failed to create product")
	ErrGetProductById   = errors.New("failed to get product by id")
	ErrProductNotFound  = errors.New("product not found")
	ErrUpdateProduct    = errors.New("no fields to update")
	ErrDeleteProduct    = errors.New("failed to delete product")
	ErrProductNameExist = errors.New("product name already exist")
)

type (
	ProductCreateRequest struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required,gt=0"`
		Stock       int     `json:"stock" validate:"gte=0"`
		ImageURL    string  `json:"image_url"`
		CategoryID  uint    `json:"category_id" validate:"required"`
	}

	ProductUpdateRequest struct {
		Name        *string  `json:"name" validate:"omitempty"`
		Description *string  `json:"description" validate:"omitempty"`
		Price       *float64 `json:"price" validate:"omitempty,gt=0"`
		Stock       *int     `json:"stock" validate:"omitempty,gte=0"`
		ImageURL    *string  `json:"image_url" validate:"omitempty"`
		CategoryID  *uint    `json:"category_id" validate:"omitempty"`
	}

	ProductResponse struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		ImageURL    string  `json:"image_url"`

		CategoryID uint `json:"category_id"`
	}
)
