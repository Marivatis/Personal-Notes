package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppEnv     string `env:"APP_ENV" env-default:"development"`
	Port       int    `env:"PORT" env-default:"8080"`
	LogFormat  string `env:"LOG_FORMAT" env-default:"console"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBSSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
