package router

import (
	"github.com/gofiber/fiber/v2"
	"possiblemeatball.com/api/v2/handler"
)

func SetupPublicRoutes(app *fiber.App) {
	api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1.0.0")
		return c.Next()
	})
	api.Get("/john/add", handler.AddJohn)
	api.Post("/john/change", handler.ChangeJohn)
}
