package repositories

import (
	"context"
	"fmt"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/models"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	redis_pkg "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type productRepository struct {
	db      *gorm.DB
	redisDb redis.RedisService
	cfg     config.Config
}

var _ ProductInterface = (*productRepository)(nil)

type NewProductRepositoryParams struct {
	Db    *gorm.DB
	Redis redis.RedisService
}

func (r *productRepository) GetProductById(ctx context.Context, p GetProductByIdParams) (*GetProductByIdResponse, error) {
	var product models.Product

	cacheKey := fmt.Sprintf("|%s|id|%s|", r.cfg.Service, p.ProductId)

	err := r.redisDb.GetCache(ctx, cacheKey, &product)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&product, "id = ? AND status <> deleted", p.ProductId).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, product, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetProductByIdResponse{
		Product: product,
	}, nil
}

func (r *productRepository) GetProductBySKU(ctx context.Context, p GetProductBySKUParams) (*GetProductBySKUResponse, error) {
	var product models.Product

	cacheKey := fmt.Sprintf("|%s|sku|%s|", r.cfg.Service, p.ProductSKU)

	err := r.redisDb.GetCache(ctx, cacheKey, &product)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&product, "sku = ? AND status <> deleted", p.ProductSKU).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, product, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetProductBySKUResponse{
		Product: product,
	}, nil
}

func (r *productRepository) GetProductBySlug(ctx context.Context, p GetProductBySlugParams) (*GetProductBySlugResponse, error) {
	var product models.Product

	cacheKey := fmt.Sprintf("|%s|slug|%s|", r.cfg.Service, p.ProductSlug)

	err := r.redisDb.GetCache(ctx, cacheKey, &product)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&product, "slug = ? AND status <> deleted", p.ProductSlug).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, product, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetProductBySlugResponse{
		Product: product,
	}, nil
}

func (r *productRepository) GetProducts(ctx context.Context, p GetProductsParams) (*GetProductsResponse, error) {
	var products []models.Product

	cacheKey := fmt.Sprintf("|%s|offset|%d|limit|%d|", r.cfg.Service, p.Offset, p.Limit)

	// todo: check redis is null or not
	err := r.redisDb.GetCache(ctx, cacheKey, &products)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		err := r.db.Model(&models.Product{}).Limit(p.Limit).Offset(p.Offset).Find(&products).Error

		if err != nil {
			return nil, err
		}

		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, products, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetProductsResponse{
		Products: products,
	}, nil
}

// func (r *userRepository) CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error) {
// 	user := models.User{
// 		Name:      p.Name,
// 		Email:     p.Email,
// 		Status:    p.Status,
// 		CreatedBy: p.CreatedBy,
// 	}

// 	err := r.db.Create(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &CreateNewUserResponse{
// 		User: user,
// 	}, nil
// }

// func (r *userRepository) UpdateUserById(ctx context.Context, p UpdateUserByIdParams) (*UpdateUserByIdResponse, error) {
// 	user := models.User{
// 		ID:        p.ID,
// 		Name:      p.Name,
// 		Email:     p.Email,
// 		Status:    p.Status,
// 		UpdatedBy: p.UpdatedBy,
// 	}

// 	err := r.db.Model(&models.User{}).Updates(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &UpdateUserByIdResponse{
// 		User: user,
// 	}, nil
// }

func (r *productRepository) DeleteProductById(ctx context.Context, p DeleteProductByIdParams) error {
	product := models.Product{
		ID: p.ProductId,
	}

	err := r.db.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func NewProductRepository(p NewProductRepositoryParams) *productRepository {

	return &productRepository{
		db:      p.Db,
		redisDb: p.Redis,
	}
}
