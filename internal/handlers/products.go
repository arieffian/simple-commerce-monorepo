package handlers

import (
	"errors"

	generated "github.com/arieffian/simple-commerces-monorepo/internal/pkg/generated/products"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type productHandler struct {
	productRepo repositories.ProductInterface
	validator   validator.ValidatorService
}

var _ ProductService = (*productHandler)(nil)

type NewProductHandlerParams struct {
	ProductRepo repositories.ProductInterface
	Validator   validator.ValidatorService
}

// // todo: add validation
// func (h *userHandler) GetUsers(c *fiber.Ctx) error {
// 	pPage := c.Params("page")
// 	page, err := strconv.Atoi(pPage)

// 	if err != nil {
// 		status := fiber.StatusBadRequest
// 		response := generated.GetUsersResponse{
// 			Code:    int32(status),
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	if page <= 0 {
// 		page = 1
// 	}

// 	usersResult, err := h.userRepo.GetUsers(c.Context(), userRepository.GetUsersParams{
// 		Limit:  PER_PAGE,
// 		Offset: (page - 1) * PER_PAGE,
// 	})

// 	if err != nil {
// 		status := fiber.StatusInternalServerError
// 		response := generated.GetUsersResponse{
// 			Code:    int32(status),
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	var users []generated.User
// 	for _, user := range usersResult.Users {
// 		users = append(users, generated.User{
// 			Id:     user.ID,
// 			Name:   user.Name,
// 			Email:  user.Email,
// 			Status: generated.UserStatus(user.Status),
// 		})
// 	}

// 	response := generated.GetUsersResponse{
// 		Code:    fiber.StatusOK,
// 		Message: "OK",
// 		Data:    &users,
// 	}

// 	return c.Status(int(response.Code)).JSON(response)
// }

// func (h *userHandler) CreateUser(c *fiber.Ctx) error {
// 	params := new(CreateNewUserParams)
// 	err := c.BodyParser(params)
// 	if err != nil {
// 		response := &generated.CreateNewUserResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	err = h.validator.Validate(params)
// 	if err != nil {
// 		response := &generated.CreateNewUserResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	user, err := h.userRepo.CreateNewUser(c.Context(), userRepository.CreateNewUserParams{
// 		Name:      params.Name,
// 		Email:     params.Email,
// 		Status:    params.Status,
// 		CreatedBy: params.CreatedBy,
// 	})

// 	if err != nil {
// 		response := &generated.CreateNewUserResponse{
// 			Code:    fiber.StatusInternalServerError,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	response := &generated.CreateNewUserResponse{
// 		Code:    fiber.StatusOK,
// 		Message: "OK",
// 		Data: &generated.User{
// 			Id:     user.User.ID,
// 			Name:   user.User.Name,
// 			Email:  user.User.Email,
// 			Status: generated.UserStatus(user.User.Status),
// 		},
// 	}

// 	return c.Status(int(response.Code)).JSON(response)
// }

// func (h *userHandler) UpdateUserById(c *fiber.Ctx) error {
// 	params := new(UpdateUserByIdParams)
// 	err := c.BodyParser(params)
// 	if err != nil {
// 		response := &generated.UpdateUserByIdResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	params.ID = c.Params("id")

// 	err = h.validator.Validate(params)
// 	if err != nil {
// 		response := &generated.UpdateUserByIdResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	user, err := h.userRepo.UpdateUserById(c.Context(), userRepository.UpdateUserByIdParams{
// 		ID:        params.ID,
// 		Name:      params.Name,
// 		Email:     params.Email,
// 		Status:    params.Status,
// 		UpdatedBy: params.UpdatedBy,
// 	})
// 	if err != nil {
// 		response := &generated.UpdateUserByIdResponse{
// 			Code:    fiber.StatusInternalServerError,
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	response := &generated.UpdateUserByIdResponse{
// 		Code:    fiber.StatusOK,
// 		Message: "OK",
// 		Data: &generated.User{
// 			Id:     user.User.ID,
// 			Name:   user.User.Name,
// 			Email:  user.User.Email,
// 			Status: generated.UserStatus(user.User.Status),
// 		},
// 	}

// 	return c.Status(int(response.Code)).JSON(response)
// }

// func (h *userHandler) DeleteUserById(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	err := h.userRepo.DeleteUserById(c.Context(), userRepository.DeleteUserByIdParams{
// 		ID: id,
// 	})

// 	if err != nil {
// 		status := fiber.StatusInternalServerError
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			status = fiber.StatusNotFound
// 		}
// 		response := generated.DeleteUserByIdResponse{
// 			Code:    int32(status),
// 			Message: err.Error(),
// 			Data:    nil,
// 		}
// 		return c.Status(int(response.Code)).JSON(response)
// 	}

// 	response := generated.DeleteUserByIdResponse{
// 		Code:    fiber.StatusOK,
// 		Message: "OK",
// 		Data:    nil,
// 	}

// 	return c.Status(int(response.Code)).JSON(response)
// }

func (h *productHandler) GetProductById(c *fiber.Ctx) error {
	product, err := h.productRepo.GetProductById(c.Context(), repositories.GetProductByIdParams{
		ProductId: c.Params("id"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetProductByIdResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetProductByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func NewProductHandler(p NewProductHandlerParams) *productHandler {
	return &productHandler{
		productRepo: p.ProductRepo,
		validator:   p.Validator,
	}
}
