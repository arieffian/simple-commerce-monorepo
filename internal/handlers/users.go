package handlers

import (
	"errors"
	"strconv"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	generated "github.com/arieffian/simple-commerces-monorepo/internal/pkg/generated/users"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userHandler struct {
	userRepo  repositories.UserInterface
	validator validator.ValidatorService
	cfg       config.Config
}

var _ UserService = (*userHandler)(nil)

type NewUserHandlerParams struct {
	UserRepo  repositories.UserInterface
	Validator validator.ValidatorService
	Cfg       config.Config
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	pPage := c.Params("page")
	page, err := strconv.Atoi(pPage)

	if err != nil {
		status := fiber.StatusBadRequest
		response := generated.GetUsersResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	if page <= 0 {
		page = 1
	}

	usersResult, err := h.userRepo.GetUsers(c.Context(), repositories.GetUsersParams{
		Limit:  PER_PAGE,
		Offset: (page - 1) * PER_PAGE,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		response := generated.GetUsersResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var users []generated.User
	for _, user := range usersResult.Users {
		users = append(users, generated.User{
			Id:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Type:   generated.UserType(user.Type),
			Status: generated.UserStatus(user.Status),
		})
	}

	response := generated.GetUsersResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    &users,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	params := new(CreateNewUserParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.CreateNewUserResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.CreateNewUserResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	user, err := h.userRepo.CreateNewUser(c.Context(), repositories.CreateNewUserParams{
		Name:      params.Name,
		Email:     params.Email,
		Status:    params.Status,
		Type:      params.Type,
		CreatedBy: params.CreatedBy,
	})

	if err != nil {
		response := &generated.CreateNewUserResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := &generated.CreateNewUserResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.User{
			Id:     user.User.ID,
			Name:   user.User.Name,
			Email:  user.User.Email,
			Type:   generated.UserType(user.User.Type),
			Status: generated.UserStatus(user.User.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *userHandler) UpdateUserById(c *fiber.Ctx) error {
	params := new(UpdateUserByIdParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.UpdateUserByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	params.ID = c.Params("id")

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.UpdateUserByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	user, err := h.userRepo.UpdateUserById(c.Context(), repositories.UpdateUserByIdParams{
		ID:        params.ID,
		Name:      params.Name,
		Email:     params.Email,
		Status:    params.Status,
		Type:      params.Type,
		UpdatedBy: params.UpdatedBy,
	})
	if err != nil {
		response := &generated.UpdateUserByIdResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := &generated.UpdateUserByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.User{
			Id:     user.User.ID,
			Name:   user.User.Name,
			Email:  user.User.Email,
			Type:   generated.UserType(user.User.Type),
			Status: generated.UserStatus(user.User.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *userHandler) DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userRepo.DeleteUserById(c.Context(), repositories.DeleteUserByIdParams{
		UserId: id,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.DeleteUserByIdResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.DeleteUserByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    nil,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
	user, err := h.userRepo.GetUserById(c.Context(), repositories.GetUserByIdParams{
		UserId: c.Params("id"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetUserByIdResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetUserByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.User{
			Id:     user.User.ID,
			Email:  user.User.Email,
			Type:   generated.UserType(user.User.Type),
			Status: generated.UserStatus(user.User.Status),
			Name:   user.User.Name,
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func NewUserHandler(p NewUserHandlerParams) *userHandler {
	return &userHandler{
		userRepo:  p.UserRepo,
		validator: p.Validator,
		cfg:       p.Cfg,
	}
}
