package handlers

import (
	"github.com/gofiber/fiber/v2"
)

const (
	PER_PAGE = 10
)

type HealthcheckService interface {
	HealthCheckHandler(c *fiber.Ctx) error
}

type UserService interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUserById(c *fiber.Ctx) error
	DeleteUserById(c *fiber.Ctx) error
}

type CreateNewUserParams struct {
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Status    string `json:"status" validate:"required,oneof=active disabled"`
	Type      string `json:"type" validate:"required,oneof=customer admin"`
	CreatedBy string `json:"created_by" validate:"required"`
}

type UpdateUserByIdParams struct {
	ID        string `json:"id" validate:"required,min=36,max=36"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Status    string `json:"status" validate:"required,oneof=active disabled"`
	Type      string `json:"type" validate:"required,oneof=customer admin"`
	UpdatedBy string `json:"updated_by" validate:"required"`
}

type ProductService interface {
	GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	GetProductBySKU(c *fiber.Ctx) error
	GetProductBySlug(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
}

type CreateNewProductParams struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	SKU         string `json:"sku" validate:"required,sku,min=6,max=255"`
	Status      string `json:"status" validate:"required,oneof=active disabled"`
	Price       uint64 `json:"price" validate:"required,min=5000"`
	Qty         uint   `json:"qty" validate:"required,min=0"`
	Description string `json:"desciption"`
	CreatedBy   string `json:"created_by" validate:"required"`
}

type UpdateProductByIdParams struct {
	ID          string `json:"id" validate:"required,min=36,max=36"`
	Name        string `json:"name" validate:"required,min=3,max=255"`
	SKU         string `json:"sku" validate:"required,sku,min=6,max=255"`
	Status      string `json:"status" validate:"required,oneof=active disabled"`
	Price       uint64 `json:"price" validate:"required,min=5000"`
	Qty         uint   `json:"qty" validate:"required,min=0"`
	Description string `json:"desciption"`
	UpdatedBy   string `json:"created_by" validate:"required"`
}

type OrderService interface {
	GetOrderById(c *fiber.Ctx) error
	GetOrdersByUserId(c *fiber.Ctx) error
	CreateNewOrder(c *fiber.Ctx) error
	UpdateOrderById(c *fiber.Ctx) error
}

type CreateNewOrderParams struct {
	UserId       string                       `json:"user_id" validate:"required,min=36,max=36"`
	OrderDetails []CreateNewOrderDetailParams `json:"details" validate:"required,min=1"`
}

type CreateNewOrderDetailParams struct {
	ProductId string `json:"product_id" validate:"required,min=36,max=36"`
	Price     uint64 `json:"price" validate:"required,min=5000"`
	Qty       uint   `json:"qty" validate:"required,min=1"`
}

type UpdateOrderByIdParams struct {
	ID     string `json:"id" validate:"required,min=36,max=36"`
	Status string `json:"status" validate:"required,oneof=pending paid canceled"`
}
