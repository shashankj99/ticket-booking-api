package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/config"
	"github.com/shashankj99/ticket-booking-api/db"
	"github.com/shashankj99/ticket-booking-api/handlers"
	"github.com/shashankj99/ticket-booking-api/repositories"
)

func main() {
	config := config.EnvConfig()
	db := db.Init(config, db.AutoMigrate)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking API",
		ServerHeader: "Ticket Booking API Header",
	})

	eventRepo := repositories.NewEventRepository(db)
	ticketRepo := repositories.NewTicketRepository(db)

	router := app.Group("/api")

	handlers.NewEventHandler(router.Group("/events"), eventRepo)
	handlers.NewTicketHandler(router.Group("/tickets"), ticketRepo)

	app.Listen(fmt.Sprintf(":%s", config.ServerPort))
}
