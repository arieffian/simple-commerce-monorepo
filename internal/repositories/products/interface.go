package users

import (
	"context"

	"github.com/arieffian/go-boilerplate/internal/models"
)

type GetProductByIdParams struct {
	ProductId string
}

type GetProductByIdResponse struct {
	Product models.Product
}

type GetProductsParams struct {
	Limit  int
	Offset int
}

type GetProductsResponse struct {
	Products []models.Product
}

type ProductInterface interface {
	GetUserById(ctx context.Context, p GetProductByIdParams) (*GetProductByIdResponse, error)
	GetUsers(ctx context.Context, p GetProductsParams) (*GetProductsResponse, error)
}
