package routers

import (
	"github.com/arieffian/simple-commerces-monorepo/internal/handlers"
	"github.com/arieffian/simple-commerces-monorepo/internal/middlewares"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type productRouter struct {
	healthcheck handlers.HealthcheckService
	products    handlers.ProductService
	apiKey      string
}

type NewProductsRouterParams struct {
	Db        *gorm.DB
	Redis     redis.RedisService
	Validator validator.ValidatorService
	APIKey    string
}

func NewProductsRouter(p NewProductsRouterParams) (*productRouter, error) {

	validator := validator.NewValidatorService()

	productRepo := repositories.NewProductRepository(repositories.NewProductRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	productHandler := handlers.NewProductHandler(handlers.NewProductHandlerParams{
		ProductRepo: productRepo,
		Validator:   validator,
	})

	return &productRouter{
		healthcheck: healthcheckHandler,
		products:    productHandler,
		apiKey:      p.APIKey,
	}, nil
}

func (r *productRouter) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1").Use(middlewares.NewValidateAPIKey(r.apiKey))
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	products := v1.Group("/products")
	// users.Get("/get-users/:page?", r.users.GetUsers)
	products.Get("/get-product-by-id/:id", r.products.GetProductById)
	// users.Post("/", r.users.CreateUser)
	// users.Patch("/:id", r.users.UpdateUserById)
	// users.Delete("/:id", r.users.DeleteUserById)
}
