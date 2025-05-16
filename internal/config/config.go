package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	Port        string `env:"PORT" envDefault:"8080"`
	GinMode     string `env:"GIN_MODE" envDefault:"debug"`
}

func LoadConfig() (*Config, error) {
	if os.Getenv("GIN_MODE") != "release" {
		godotenv.Load()
	}

	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
