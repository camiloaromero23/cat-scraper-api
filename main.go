package main

import (
	"fmt"
	"log"

	"github.com/camiloaromero23/cat-scraper-api/handlers"
	"github.com/camiloaromero23/cat-scraper-api/middlewares"
	"github.com/camiloaromero23/cat-scraper-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
  host := utils.GoDotEnvVariable("HOST")
  port := utils.GoDotEnvVariable("PORT")

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

  server := fmt.Sprintf("%s:%s", host, port)

	log.Fatal(app.Listen(server))
}
