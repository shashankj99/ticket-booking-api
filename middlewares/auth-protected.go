package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shashankj99/ticket-booking-api/models"
	"github.com/shashankj99/ticket-booking-api/utils"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			log.Warnf("No token provided")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Warnf("Invalid token format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		parsedToken, err := utils.ValidateToken(tokenParts[1], os.Getenv("JWT_SECRET"))
		if err != nil {
			log.Warnf("Invalid token: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		userID := parsedToken.Claims.(jwt.MapClaims)["id"].(string)

		var user models.User
		if err := db.Model(&models.User{}).Where("id = ?", userID).First(&user).Error; err != nil {
			log.Warnf("User not found: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
