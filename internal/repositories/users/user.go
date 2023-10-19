package users

import (
	"context"
	"strconv"

	"github.com/arieffian/go-boilerplate/internal/models"
	"github.com/arieffian/go-boilerplate/internal/pkg/redis"
	redis_pkg "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRepository struct {
	db      *gorm.DB
	redisDb redis.RedisService
}

var _ UserInterface = (*userRepository)(nil)

type NewUserRepositoryParams struct {
	Db    *gorm.DB
	Redis redis.RedisService
}

func (r *userRepository) GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error) {
	var user models.User

	cacheKey := "|user|id|" + p.UserId + "|"

	err := r.redisDb.GetCache(ctx, cacheKey, &user)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&user, "id = ? AND status <> deleted", p.UserId).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, user, 0); err != nil {
			return nil, err
		}
	}

	return &GetUserByIdResponse{
		User: user,
	}, nil
}

func (r *userRepository) GetUsers(ctx context.Context, p GetUsersParams) (*GetUsersResponse, error) {
	var users []models.User

	cacheKey := "|users|offset|" + strconv.Itoa(p.Offset) + "|limit|" + strconv.Itoa(p.Limit) + "|"

	err := r.redisDb.GetCache(ctx, cacheKey, &users)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		err := r.db.Model(&models.User{}).Limit(p.Limit).Offset(p.Offset).Find(&users).Error

		if err != nil {
			return nil, err
		}

		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, users, 0); err != nil {
			return nil, err
		}
	}

	return &GetUsersResponse{
		Users: users,
	}, nil
}

func (r *userRepository) CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error) {
	user := models.User{
		Name:      p.Name,
		Email:     p.Email,
		Status:    p.Status,
		CreatedBy: p.CreatedBy,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &CreateNewUserResponse{
		User: user,
	}, nil
}

func (r *userRepository) UpdateUserById(ctx context.Context, p UpdateUserByIdParams) (*UpdateUserByIdResponse, error) {
	user := models.User{
		ID:        p.ID,
		Name:      p.Name,
		Email:     p.Email,
		Status:    p.Status,
		UpdatedBy: p.UpdatedBy,
	}

	err := r.db.Model(&models.User{}).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &UpdateUserByIdResponse{
		User: user,
	}, nil
}

func (r *userRepository) DeleteUserById(ctx context.Context, p DeleteUserByIdParams) error {
	user := models.User{
		ID: p.ID,
	}

	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(p NewUserRepositoryParams) *userRepository {

	return &userRepository{
		db:      p.Db,
		redisDb: p.Redis,
	}
}
