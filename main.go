package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("views/index.html")
	})

	app.Get("/partial", func(c *fiber.Ctx) error {
		return c.SendFile("views/partial.html")
	})

	app.Listen(":3000")
}
