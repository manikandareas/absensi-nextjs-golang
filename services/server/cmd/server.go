package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"code":    fiber.StatusOK,
			"message": "Api Successfully running",
		})
	})
	log.Fatal(app.Listen(":8000"))
}
