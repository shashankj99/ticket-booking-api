package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/models"
)

type EventHandler struct {
	repo models.EventRepository
}

func (h *EventHandler) FindMany(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) FindOne(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) Create(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) Delete(c *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repo models.EventRepository) {
	handler := &EventHandler{repo: repo}

	router.Get("/", handler.FindMany)
	router.Get("/:id", handler.FindOne)
	router.Post("/", handler.Create)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}
