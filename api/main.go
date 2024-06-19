package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"possiblemeatball.com/api/v2/config"
)

type Name struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	appConfig := config.EnvToFiber()
	app := fiber.New(appConfig)
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrateErr := db.AutoMigrate(&Name{})
	if migrateErr != nil {
		return
	}

	db.Create(&Name{FirstName: "John", LastName: "Doe"})

	api := app.Group("/api/v2", func(c *fiber.Ctx) error {
		c.Set("Version", "2.0.0")
		return c.Next()
	})

	api.Post("/ballsack", func(c *fiber.Ctx) error {
		var john Name
		db.First(&john, "first_name = ?", "John")
		db.Model(&john).Update("LastName", "Deere")
		return c.JSON(&john)
	})

	log.Fatal(app.Listen(":3000"))
}
