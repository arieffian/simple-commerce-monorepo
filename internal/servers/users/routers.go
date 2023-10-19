package users

import (
	healthcheckHandlers "github.com/arieffian/go-boilerplate/internal/handlers/healthcheck"
	usersHandlers "github.com/arieffian/go-boilerplate/internal/handlers/users"
	"github.com/arieffian/go-boilerplate/internal/middlewares"
	"github.com/arieffian/go-boilerplate/internal/pkg/redis"
	"github.com/arieffian/go-boilerplate/internal/pkg/validator"
	userRepository "github.com/arieffian/go-boilerplate/internal/repositories/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router struct {
	healthcheck healthcheckHandlers.HealthcheckService
	users       usersHandlers.UserService
	apiKey      string
}

type NewRouterParams struct {
	Db        *gorm.DB
	Redis     redis.RedisService
	Validator validator.ValidatorService
	APIKey    string
}

func NewRouter(p NewRouterParams) (*Router, error) {

	validator := validator.NewValidatorService()

	userRepo := userRepository.NewUserRepository(userRepository.NewUserRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
	})

	healthcheckHandler := healthcheckHandlers.NewHealthCheckHandler()
	userHandler := usersHandlers.NewUserHandler(usersHandlers.NewUserHandlerParams{
		UserRepo:  userRepo,
		Validator: validator,
	})

	return &Router{
		healthcheck: healthcheckHandler,
		users:       userHandler,
		apiKey:      p.APIKey,
	}, nil
}

func (r *Router) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1").Use(middlewares.NewValidateAPIKey(r.apiKey))
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	users := v1.Group("/users")
	users.Get("/get-users/:page?", r.users.GetUsers)
	users.Get("/get-user-by-id/:id", r.users.GetUserById)
	users.Post("/", r.users.CreateUser)
	users.Patch("/:id", r.users.UpdateUserById)
	users.Delete("/:id", r.users.DeleteUserById)
}
