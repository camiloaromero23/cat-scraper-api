package handlers

import (
	"net/http"

	catModel "github.com/camiloaromero23/cat-scraper-api/models/cat"
	"github.com/gofiber/fiber/v2"
)

func GetCatHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	cat, err := catModel.GetCat(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	if cat == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cat not found",
		})
	}

	return c.JSON(cat)
}

func GetCatsHandler(c *fiber.Ctx) error {
	cats, err := catModel.GetCats()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.JSON(cats)
}

func UpdateCatsHandler(c *fiber.Ctx) error {
	cats, err := catModel.UpdateCats()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.JSON(cats)
}
