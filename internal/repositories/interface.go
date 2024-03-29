package repositories

import (
	"context"

	"github.com/arieffian/simple-commerces-monorepo/internal/models"
)

type GetUserByIdParams struct {
	UserId string
}

type GetUserByIdResponse struct {
	User models.User
}

type GetUsersParams struct {
	Limit  int
	Offset int
}

type GetUsersResponse struct {
	Users []models.User
}

type CreateNewUserParams struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	CreatedBy string `json:"created_by"`
}

type CreateNewUserResponse struct {
	User models.User
}

type UpdateUserByIdParams struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateUserByIdResponse struct {
	User models.User
}

type DeleteUserByIdParams struct {
	UserId string `json:"user_id"`
}

type UserInterface interface {
	GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error)
	GetUsers(ctx context.Context, p GetUsersParams) (*GetUsersResponse, error)
	CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error)
	UpdateUserById(ctx context.Context, p UpdateUserByIdParams) (*UpdateUserByIdResponse, error)
	DeleteUserById(ctx context.Context, p DeleteUserByIdParams) error
}

type GetProductByIdParams struct {
	ProductId string
}

type GetProductByIdResponse struct {
	Product models.Product
}

type GetProductBySlugParams struct {
	ProductSlug string
}

type GetProductBySlugResponse struct {
	Product models.Product
}

type GetProductBySKUParams struct {
	ProductSKU string
}

type GetProductBySKUResponse struct {
	Product models.Product
}

type GetProductsParams struct {
	Limit  int
	Offset int
}

type GetProductsResponse struct {
	Products []models.Product
}

type CreateNewProductParams struct {
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Price       uint64 `json:"price"`
	Qty         uint   `json:"qty"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
}

type CreateNewProductResponse struct {
	Product models.Product
}

type UpdateProductByIdParams struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Price       uint64 `json:"price"`
	Qty         uint   `json:"qty"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updated_by"`
}

type UpdateProductByIdResponse struct {
	Product models.Product
}

type DeleteProductByIdParams struct {
	ProductId string `json:"product_id"`
}

type ProductInterface interface {
	GetProductById(ctx context.Context, p GetProductByIdParams) (*GetProductByIdResponse, error)
	GetProductBySKU(ctx context.Context, p GetProductBySKUParams) (*GetProductBySKUResponse, error)
	GetProductBySlug(ctx context.Context, p GetProductBySlugParams) (*GetProductBySlugResponse, error)
	GetProducts(ctx context.Context, p GetProductsParams) (*GetProductsResponse, error)
	CreateNewProduct(ctx context.Context, p CreateNewProductParams) (*CreateNewProductResponse, error)
	UpdateProductById(ctx context.Context, p UpdateProductByIdParams) (*UpdateProductByIdResponse, error)
	DeleteProductById(ctx context.Context, p DeleteProductByIdParams) error
}

type GetOrderByIdParams struct {
	OrderId string
}

type GetOrderByIdResponse struct {
	Order models.Order
}

type GetOrdersByUserIdParams struct {
	UserId string
	Limit  int
	Offset int
}

type GetOrdersByUserIdResponse struct {
	Orders []models.Order
}

type CreateNewOrderDetailParams struct {
	ProductId string `json:"product_id"`
	SubTotal  uint64 `json:"sub_total"`
	Qty       uint   `json:"qty"`
	Price     uint32 `json:"price"`
}

type CreateNewOrderParams struct {
	UserId       string                       `json:"user_id"`
	OrderDetails []CreateNewOrderDetailParams `json:"order_details"`
	GrandTotal   uint64                       `json:"grand_total"`
	Status       string                       `json:"status"`
}

type CreateNewOrderResponse struct {
	Order models.Order
}

type UpdateOrderByIdParams struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateOrderByIdResponse struct {
	Order models.Order
}

type OrderInterface interface {
	GetOrderById(ctx context.Context, p GetOrderByIdParams) (*GetOrderByIdResponse, error)
	CreateNewOrder(ctx context.Context, p CreateNewOrderParams) (*CreateNewOrderResponse, error)
	GetOrdersByUserId(ctx context.Context, p GetOrdersByUserIdParams) (*GetOrdersByUserIdResponse, error)
	UpdateOrderById(ctx context.Context, p UpdateOrderByIdParams) (*UpdateOrderByIdResponse, error)
}
