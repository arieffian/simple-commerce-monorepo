package config

import (
	"context"
	"errors"
	"os"

	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Environment               string `mapstructure:"ENVIRONMENT" validate:"required"`
	Debug                     bool   `mapstructure:"DEBUG" validate:"required"`
	DbMasterConnectionString  string `mapstructure:"DB_MASTER_CONNECTION_STRING" validate:"required"`
	DbReplicaConnectionString string `mapstructure:"DB_REPLICA_CONNECTION_STRING"`
	APIAddress                string `mapstructure:"API_ADDRESS" validate:"required"`
	RedisHost                 string `mapstructure:"REDIS_HOST"`
	RedisPort                 int    `mapstructure:"REDIS_PORT"`
	CacheTTL                  int    `mapstructure:"CACHE_TTL"`
	APIKey                    string `mapstructure:"API_KEY" validate:"required"`
	DBDriver                  string `mapstructure:"DB_DRIVER" validate:"required"`
	Service                   string `mapstructure:"SERVICE" validate:"required"`
}

func NewConfig() (*Config, error) {
	serviceType := os.Getenv("SERVICE_TYPE")

	if serviceType == "" {
		log.Printf("service type cannot be empty")
		return nil, errors.New("service type cannot be empty")
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName(serviceType + ".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Could not load config with error: %s", err.Error())
		return nil, err
	}

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Printf("Failed to load env variables. %+v\n", err)
		return nil, err
	}

	v := validator.NewValidatorService()
	if err := v.Validate(context.Background(), cfg); err != nil {
		log.Printf("Failed to validate config. %+v\n", err)
		return nil, err
	}

	return &cfg, nil
}
