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

type ProductService interface {
	// GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	// CreateProduct(c *fiber.Ctx) error
	// UpdateProductById(c *fiber.Ctx) error
	// DeleteProductById(c *fiber.Ctx) error
}

type CreateNewUserParams struct {
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Status    string `json:"status" validate:"required,oneof=active disabled"`
	CreatedBy string `json:"created_by" validate:"required"`
}

type UpdateUserByIdParams struct {
	ID        string `json:"id" validate:"required,min=36,max=36"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Email     string `json:"email" validate:"required,email"`
	Status    string `json:"status" validate:"required,oneof=active disabled"`
	UpdatedBy string `json:"updated_by" validate:"required"`
}