package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type IConfig interface {
	GetEnvConfig()
	GetDatabase() *Database
	GetService() *Service
}

type Configuration struct {
	Database *Database
	Service  *Service
}

type Database struct {
	Host     string `json:"host" env:"DB_HOST" default:"localhost"`
	Port     string `json:"port" env:"DB_PORT" default:"5432"`
	Username string `json:"username" env:"DB_USERNAME" default:"postgres"`
	Password string `json:"password" env:"DB_PASSWORD" default:"postgres"`
	Name     string `json:"name" env:"DB_NAME" default:"postgres"`
	Scheme   string `json:"scheme" env:"DB_SCHEME" default:"postgres"`
	Tcp      string `json:"tcp" env:"DB_TCP" default:"tcp"`
}

type Service struct {
	Port string `json:"port" env:"SERVICE_PORT" default:"8080"`
	Host string `json:"host" env:"SERVICE_HOST" default:"localhost"`
}

func New() IConfig {
	return &Configuration{}
}

func (config *Configuration) GetEnvConfig() {
	ctx := context.Background()
	if err := envconfig.Process(ctx, config); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}
}

func (config *Configuration) GetDatabase() *Database {
	return config.Database
}

func (config *Configuration) GetService() *Service {
	return config.Service
}
