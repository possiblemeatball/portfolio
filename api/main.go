package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v3/log"
	"possiblemeatball.com/api/v2/config"
	"possiblemeatball.com/api/v2/router"
)

func main() {
	appConfig := config.EnvToFiber()
	app := fiber.New(appConfig)
	app.Use(cors.New())
	app.Use(logger.New())
	router.SetupPublicRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
