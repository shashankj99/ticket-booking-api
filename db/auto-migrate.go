package db

import (
	"github.com/shashankj99/ticket-booking-api/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
