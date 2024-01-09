package handler

import (
	"github.com/ashutoshbr/Fiber-GORM/database"
	"github.com/ashutoshbr/Fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(&users)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	database.DB.Create(&user)
	return c.JSON(&user)
}
