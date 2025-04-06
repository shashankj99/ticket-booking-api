package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/handlers"
	"github.com/shashankj99/ticket-booking-api/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking API",
		ServerHeader: "Ticket Booking API Header",
	})

	// Repositories
	eventRepo := repositories.NewEventRepository(nil)

	// Routing
	router := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(router.Group("/events"), eventRepo)

	app.Listen(":3000")
}
