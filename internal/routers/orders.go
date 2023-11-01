package routers

import (
	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	"github.com/arieffian/simple-commerces-monorepo/internal/handlers"
	"github.com/arieffian/simple-commerces-monorepo/internal/middlewares"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type orderRouter struct {
	healthcheck handlers.HealthcheckService
	orders      handlers.OrderService
	cfg         config.Config
}

type NewOrdersRouterParams struct {
	Db    *database.DbInstance
	Redis redis.RedisService
	Cfg   config.Config
}

func NewOrdersRouter(p NewUsersRouterParams) (*orderRouter, error) {

	validator := validator.NewValidatorService()

	orderRepo := repositories.NewOrderRepository(repositories.NewOrderRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
		Cfg:   p.Cfg,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	orderHandler := handlers.NewOrderHandler(handlers.NewOrderHandlerParams{
		OrderRepo: orderRepo,
		Validator: validator,
		Cfg:       p.Cfg,
	})

	return &orderRouter{
		healthcheck: healthcheckHandler,
		orders:      orderHandler,
		cfg:         p.Cfg,
	}, nil
}

func (r *orderRouter) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1").Use(middlewares.NewValidateAPIKey(r.cfg.APIKey))
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	orders := v1.Group("/orders")
	orders.Get("/get-orders-by-user-id/:user_id/:page?", r.orders.GetOrdersByUserId)
	orders.Get("/get-order-by-id/:id", r.orders.GetOrderById)
	orders.Post("/", r.orders.CreateNewOrder)
	orders.Patch("/:id", r.orders.UpdateOrderById)
}
