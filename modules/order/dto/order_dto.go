package dto

import "errors"

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_SUCCESS_GET_DATA          = "success get data"
)

const (
	MESSAGE_FAILED_CREATE_ORDER   = "failed create order"
	MESSAGE_FAILED_GET_ORDER      = "failed get order"
	MESSAGE_FAILED_GET_LIST_ORDER = "failed get order list"
	MESSAGE_FAILED_DELETE_ORDER   = "failed delete order"
	MESSAGE_FAILED_PROSES_REQUEST = "failed proses request"

	// Success
	MESSAGE_SUCCESS_CREATE_PRODUCT   = "success create order"
	MESSAGE_SUCCESS_GET_PRODUCT      = "success get order"
	MESSAGE_SUCCESS_GET_LIST_PRODUCT = "success get order list"
	MESSAGE_SUCCESS_DELETE_PRODUCT   = "success delete order"
)

var (
	ErrCreateOrder    = errors.New("failed to create order")
	ErrGetOrderById   = errors.New("failed to get order by id")
	ErrorderNotFound  = errors.New("order not found")
	ErrDeleteOrder    = errors.New("failed to delete order")
	ErrorderNameExist = errors.New("order name already exist")
)

type (
	OrderCreateRequest struct {
	}

	OrderResponse struct {
	}
)
