package main

import (
	"zanuarmirza/goFiberSample/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/copy/:foo", handler.CopyHandler)
	app.Post("/write", handler.WriteHandler)

	app.Listen(":3000")
}
