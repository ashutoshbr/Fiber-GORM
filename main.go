package main

import (
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Namaste")
}
func main() {
	app := fiber.New()
	app.Get("/", home)
	app.Listen(":8000")
}
