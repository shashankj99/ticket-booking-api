package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/models"
)

type TicketHandler struct {
	repo models.TicketRepository
}

func (h *TicketHandler) FindMany(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickets, err := h.repo.FindMany(context)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(tickets)
}

func (h *TicketHandler) FindOne(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticket, err := h.repo.FindOne(context, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(ticket)
}

func (h *TicketHandler) Create(c *fiber.Ctx) error {
	ticket := &models.Ticket{}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.BodyParser(ticket); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	ticket, err := h.repo.Create(context, ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(ticket)
}

func (h *TicketHandler) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	updateData := make(map[string]any)

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	ticket, err := h.repo.Update(context, uint(id), updateData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(ticket)
}

func (h *TicketHandler) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.repo.Delete(context, uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func NewTicketHandler(router fiber.Router, repo models.TicketRepository) {
	handler := &TicketHandler{repo: repo}

	router.Get("/", handler.FindMany)
	router.Get("/:id", handler.FindOne)
	router.Post("/", handler.Create)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}
