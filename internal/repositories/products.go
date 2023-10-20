package repositories

import (
	"context"
	"fmt"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	"github.com/arieffian/simple-commerces-monorepo/internal/models"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/slug"
	redis_pkg "github.com/redis/go-redis/v9"
)

type productRepository struct {
	db            *database.DbInstance
	redisDb       redis.RedisService
	cfg           config.Config
	slugGenerator slug.SlugGeneratorService
}

var _ ProductInterface = (*productRepository)(nil)

type NewProductRepositoryParams struct {
	Db            *database.DbInstance
	Redis         redis.RedisService
	Cfg           config.Config
	SlugGenerator slug.SlugGeneratorService
}

func (r *productRepository) GetProductById(ctx context.Context, p GetProductByIdParams) (*GetProductByIdResponse, error) {
	var product models.Product

	cacheKey := fmt.Sprintf("|%s|id|%s|", r.cfg.Service, p.ProductId)

	err := r.redisDb.GetCache(ctx, cacheKey, &product)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.Db.First(&product, "id = ? AND status <> deleted", p.ProductId).Error; err != nil {
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

		if err := r.db.Db.First(&product, "sku = ? AND status <> deleted", p.ProductSKU).Error; err != nil {
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

		if err := r.db.Db.First(&product, "slug = ? AND status <> deleted", p.ProductSlug).Error; err != nil {
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

		err := r.db.Db.Model(&models.Product{}).Limit(p.Limit).Offset(p.Offset).Find(&products).Error

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

func (r *productRepository) CreateNewProduct(ctx context.Context, p CreateNewProductParams) (*CreateNewProductResponse, error) {
	product := models.Product{
		Name:        p.Name,
		SKU:         p.SKU,
		Price:       p.Price,
		Qty:         p.Qty,
		Status:      p.Status,
		Description: p.Description,
		CreatedBy:   p.CreatedBy,
	}

	slug, err := r.slugGenerator.GenerateUniqueSlug(ctx, p.Name, "products", "slug")
	if err != nil {
		return nil, err
	}

	product.Slug = slug

	err = r.db.Db.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &CreateNewProductResponse{
		Product: product,
	}, nil
}

func (r *productRepository) UpdateProductById(ctx context.Context, p UpdateProductByIdParams) (*UpdateProductByIdResponse, error) {
	product := models.Product{
		Name:        p.Name,
		SKU:         p.SKU,
		Price:       p.Price,
		Qty:         p.Qty,
		Status:      p.Status,
		Description: p.Description,
		UpdatedBy:   p.UpdatedBy,
	}

	err := r.db.Db.Model(&models.Product{}).Updates(&product).Error
	if err != nil {
		return nil, err
	}

	return &UpdateProductByIdResponse{
		Product: product,
	}, nil
}

func (r *productRepository) DeleteProductById(ctx context.Context, p DeleteProductByIdParams) error {
	product := models.Product{
		ID: p.ProductId,
	}

	err := r.db.Db.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func NewProductRepository(p NewProductRepositoryParams) *productRepository {

	return &productRepository{
		db:            p.Db,
		redisDb:       p.Redis,
		cfg:           p.Cfg,
		slugGenerator: p.SlugGenerator,
	}
}
