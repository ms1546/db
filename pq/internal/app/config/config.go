package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Address string `envconfig:"SERVER_ADDRESS" default:":8080"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
	Username string `envconfig:"DB_USER" default:"user"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
	DBName   string `envconfig:"DB_NAME" default:"mydb"`
}

func Load() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to parse env vars: %s", err)
	}
	return cfg
}
