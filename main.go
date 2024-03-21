package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, meatball!")
	})

	log.Fatal(app.Listen(":3000"))
}
