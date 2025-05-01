package handlers

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/shashankj99/ticket-booking-api/models"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func NewAuthHandler(router fiber.Router, service models.AuthService) {
	handler := &AuthHandler{service: service}
	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var registerData models.AuthCredentials
	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(registerData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, user, err := h.service.Register(c.Context(), &registerData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
		"token":   token,
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginData models.AuthCredentials

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, user, err := h.service.Login(context, &loginData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
		"token":   token,
		"user":    user,
	})
}
