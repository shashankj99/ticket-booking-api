package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE,required"`
}

func EnvConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	return cfg
}
