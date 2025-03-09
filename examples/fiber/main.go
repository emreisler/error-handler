package main

import (
	"github.com/emreisler/error-handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(error_handler.FiberMiddleware())

	app.Get("/example", func(c *fiber.Ctx) error {
		return error_handler.BadRequestError("Invalid request")
	})

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
