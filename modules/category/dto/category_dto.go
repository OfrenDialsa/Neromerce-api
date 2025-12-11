package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_CATEGORY    = "failed create category"
	MESSAGE_FAILED_GET_CATEGORY       = "failed get category"
	MESSAGE_FAILED_GET_LIST_CATEGORY  = "failed get list category"
	MESSAGE_FAILED_DELETE_CATEGORY    = "failed delete category"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"

	// Success
	MESSAGE_SUCCESS_CREATE_CATEGORY   = "success create category"
	MESSAGE_SUCCESS_GET_CATEGORY      = "success get category"
	MESSAGE_SUCCESS_GET_LIST_CATEGORY = "success get list category"
	MESSAGE_SUCCESS_DELETE_CATEGORY   = "success delete category"
)

var (
	ErrCreateCategory    = errors.New("failed to create category")
	ErrGetCategoryById   = errors.New("failed to get category by id")
	ErrCategoryNotFound  = errors.New("category not found")
	ErrDeleteCategory    = errors.New("failed to delete category")
	ErrCategoryNameExist = errors.New("category name already exist")
)

type (
	CategoryCreateRequest struct {
		Name string `json:"name" form:"name" binding:"required,min=2,max=100"`
	}

	CategoryResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
