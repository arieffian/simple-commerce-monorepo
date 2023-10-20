package routers

import (
	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	"github.com/arieffian/simple-commerces-monorepo/internal/handlers"
	"github.com/arieffian/simple-commerces-monorepo/internal/middlewares"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/slug"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type productRouter struct {
	healthcheck handlers.HealthcheckService
	products    handlers.ProductService
	cfg         config.Config
}

type NewProductsRouterParams struct {
	Db    *database.DbInstance
	Redis redis.RedisService
	Cfg   config.Config
}

func NewProductsRouter(p NewProductsRouterParams) (*productRouter, error) {

	validator := validator.NewValidatorService()
	slugGenerator := slug.NewSlugGeneratorService(slug.NewSlugGeneratorParams{
		Db: p.Db,
	})

	productRepo := repositories.NewProductRepository(repositories.NewProductRepositoryParams{
		Db:            p.Db,
		Redis:         p.Redis,
		Cfg:           p.Cfg,
		SlugGenerator: slugGenerator,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	productHandler := handlers.NewProductHandler(handlers.NewProductHandlerParams{
		ProductRepo: productRepo,
		Validator:   validator,
		Cfg:         p.Cfg,
	})

	return &productRouter{
		healthcheck: healthcheckHandler,
		products:    productHandler,
		cfg:         p.Cfg,
	}, nil
}

func (r *productRouter) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1").Use(middlewares.NewValidateAPIKey(r.cfg.APIKey))
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	products := v1.Group("/products")
	products.Get("/get-products/:page?", r.products.GetProducts)
	products.Get("/get-product-by-id/:id", r.products.GetProductById)
	products.Get("/get-product-by-slug/:slug", r.products.GetProductBySlug)
	products.Get("/get-product-by-sku/:sku", r.products.GetProductBySKU)
	products.Post("/", r.products.CreateProduct)
	products.Patch("/:id", r.products.UpdateProductById)
	products.Delete("/:id", r.products.DeleteProductById)
}
