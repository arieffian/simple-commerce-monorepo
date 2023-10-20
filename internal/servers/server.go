package servers

import (
	"context"
	"errors"
	"log"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	"github.com/arieffian/simple-commerces-monorepo/internal/routers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber *fiber.App
}

func NewServer(ctx context.Context, cfg config.Config) (*Server, error) {

	db := database.NewDbManager(database.DbConfig{
		WriteDsn: cfg.DbMasterConnectionString,
		ReadDsn:  cfg.DbReplicaConnectionString,
	})

	var dbClient *database.DbInstance
	switch cfg.DBDriver {
	case "sqlite":
		client, err := db.CreateDbSqliteClient(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to database. %+v", err)
			return nil, err
		}
		dbClient = client
	case "postgres":
		client, err := db.CreateDbPostgresClient(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to database. %+v", err)
			return nil, err
		}
		dbClient = client
	case "mysql":
		client, err := db.CreateDbMysqlClient(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to database. %+v", err)
			return nil, err
		}
		dbClient = client
	default:
		log.Fatalf("Failed to connect to database. empty driver")
		return nil, errors.New("Failed to connect to database. empty driver")
	}

	// todo: check config cache is enabled or not
	redis := redis.NewRedisConnection(redis.RedisConfig{
		Host: cfg.RedisHost,
		Port: cfg.RedisPort,
	})

	app := fiber.New()

	var api routers.RouterService
	switch cfg.Service {
	case "users":
		router, err := routers.NewUsersRouter(routers.NewUsersRouterParams{
			Db:    dbClient,
			Redis: redis,
			Cfg:   cfg,
		})
		if err != nil {
			return nil, err
		}
		api = router
	case "products":
		router, err := routers.NewProductsRouter(routers.NewProductsRouterParams{
			Db:    dbClient,
			Redis: redis,
			Cfg:   cfg,
		})
		if err != nil {
			return nil, err
		}
		api = router
	default:
		return nil, errors.New("invalid service")
	}

	api.RegisterRoutes(app)

	return &Server{
		Fiber: app,
	}, nil

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Fiber.ShutdownWithContext(ctx)
}

func (s *Server) Listen(addr string) error {
	return s.Fiber.Listen(addr)
}
