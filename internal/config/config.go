package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string

	PGConfig PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
	Timezone string
}

func New() *AppConfig {
	_ = godotenv.Load()

	return &AppConfig{
		Port: os.Getenv("PORT"),
		PGConfig: PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Timezone: os.Getenv("DB_TIMEZONE"),
		},
	}
}

