package router

import (
	"github.com/ashutoshbr/Fiber-GORM/handler"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/users", handler.GetUsers)
	app.Post("/user", handler.SaveUser)
}
