package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/camiloaromero23/cat-scraper-api/db"
	"github.com/camiloaromero23/cat-scraper-api/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthMiddleware(c *fiber.Ctx) error {
	db := db.GetDB()

	if err := db.AutoMigrate(&types.API_KEY{}); err != nil {
		log.Println("Error migrating api_key db", err)
	}

	apiKey := c.Get("Authorization")

	var dbApiKey types.API_KEY
	res := db.First(&dbApiKey, "api_key = ?", apiKey)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting api key",
			"error":   res.Error,
		})
	}

	return c.Next()
}

