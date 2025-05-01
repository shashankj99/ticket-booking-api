package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/config"
	"github.com/shashankj99/ticket-booking-api/db"
	"github.com/shashankj99/ticket-booking-api/handlers"
	"github.com/shashankj99/ticket-booking-api/middlewares"
	"github.com/shashankj99/ticket-booking-api/repositories"
	"github.com/shashankj99/ticket-booking-api/services"
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
	authRepo := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepo)

	router := app.Group("/api")
	handlers.NewAuthHandler(router.Group("/auth"), authService)

	privateRoutes := app.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/events"), eventRepo)
	handlers.NewTicketHandler(privateRoutes.Group("/tickets"), ticketRepo)

	app.Listen(fmt.Sprintf(":%s", config.ServerPort))
}
