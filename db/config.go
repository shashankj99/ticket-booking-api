package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/shashankj99/ticket-booking-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.Config, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Info("Connected to the database!")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
