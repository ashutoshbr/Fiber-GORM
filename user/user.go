package user

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func InitialMigration() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "testing."}})
	if err != nil {
		fmt.Println(err.Error())
		panic("Error! Connecting to DB")
	}
	db.AutoMigrate()
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	db.First(&users)
	return c.JSON(&users)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Create(&user)
	return c.JSON(&user)
}
