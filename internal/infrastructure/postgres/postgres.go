package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"test-api/internal/config"
	"test-api/internal/domain/todo"
	"test-api/internal/domain/user"
	)

func Connect(cfg *config.PostgresConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DB, cfg.Port, cfg.SSLMode, cfg.Timezone)

	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		panic("Error connecting to database")
	}

	log.Println("Connected to database")

    db.AutoMigrate(
    &todo.todo{},
    &user.user{},
    )

	return db
}
