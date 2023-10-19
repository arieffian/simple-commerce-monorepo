package users

import (
	"context"

	"github.com/arieffian/go-boilerplate/internal/models"
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
	UpdatedBy string `json:"updated_by"`
}

type UpdateUserByIdResponse struct {
	User models.User
}

type DeleteUserByIdParams struct {
	ID string `json:"id"`
}

type UserInterface interface {
	GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error)
	GetUsers(ctx context.Context, p GetUsersParams) (*GetUsersResponse, error)
	CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error)
	UpdateUserById(ctx context.Context, p UpdateUserByIdParams) (*UpdateUserByIdResponse, error)
	DeleteUserById(ctx context.Context, p DeleteUserByIdParams) error
}
