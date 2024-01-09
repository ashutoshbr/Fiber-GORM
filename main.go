package main

import (
	"github.com/ashutoshbr/Fiber-GORM/database"
	"github.com/ashutoshbr/Fiber-GORM/router"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Namaste")
}

func main() {
	database.InitialMigration()
	app := fiber.New()
	app.Get("/", Home)
	router.Routers(app)
	app.Listen(":8000")
}
