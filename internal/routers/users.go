package routers

import (
	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/handlers"
	"github.com/arieffian/simple-commerces-monorepo/internal/middlewares"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userRouter struct {
	healthcheck handlers.HealthcheckService
	users       handlers.UserService
	cfg         config.Config
}

type NewUsersRouterParams struct {
	Db    *gorm.DB
	Redis redis.RedisService
	Cfg   config.Config
}

func NewUsersRouter(p NewUsersRouterParams) (*userRouter, error) {

	validator := validator.NewValidatorService()

	userRepo := repositories.NewUserRepository(repositories.NewUserRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	userHandler := handlers.NewUserHandler(handlers.NewUserHandlerParams{
		UserRepo:  userRepo,
		Validator: validator,
	})

	return &userRouter{
		healthcheck: healthcheckHandler,
		users:       userHandler,
		cfg:         p.Cfg,
	}, nil
}

func (r *userRouter) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1").Use(middlewares.NewValidateAPIKey(r.cfg.APIKey))
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	users := v1.Group("/users")
	users.Get("/get-users/:page?", r.users.GetUsers)
	users.Get("/get-user-by-id/:id", r.users.GetUserById)
	users.Post("/", r.users.CreateUser)
	users.Patch("/:id", r.users.UpdateUserById)
	users.Delete("/:id", r.users.DeleteUserById)
}
