package main

import (
	"log"

	"github.com/camiloaromero23/cat-scraper-api/handlers"
	"github.com/camiloaromero23/cat-scraper-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Use(middlewares.AuthMiddleware)

	api := app.Group("/api")

	api.Get("/cats", handlers.GetCatsHandler)
	api.Get("/cats/:id", handlers.GetCatHandler)
	api.Put("/cats", handlers.UpdateCatsHandler)

	log.Fatal(app.Listen(":3000"))
}
