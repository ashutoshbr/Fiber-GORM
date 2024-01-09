package main

import (
	"github.com/ashutoshbr/Fiber-GORM/user"
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Namaste")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Post("/user", user.SaveUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", home)
	Routers(app)
	app.Listen(":8000")
}
