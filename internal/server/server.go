package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"test-api/internal/config"
	"test-api/internal/infrastructure/postgres"
)

type App struct {
	*fiber.App
	DB *gorm.DB
}

func New() *App {
	cfg := config.New()

	db := postgres.Connect(&cfg.PGConfig)

	s := &App{
		App: fiber.New(fiber.Config{
			AppName:      "test-api",
		}),
		DB: db,
	}

	return s
}
