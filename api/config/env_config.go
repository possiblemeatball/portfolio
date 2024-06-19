package config

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"time"
)

func EnvToFiber() fiber.Config {
	prefork, _ := strconv.ParseBool(os.Getenv("FIBER_PREFORK"))
	caseSensitive, _ := strconv.ParseBool(os.Getenv("FIBER_CASE_SENSITIVE"))
	strictRouting, _ := strconv.ParseBool(os.Getenv("FIBER_STRICT_ROUTING"))
	readTimeout, _ := strconv.ParseInt(os.Getenv("FIBER_READ_TIMEOUT"), 10, 64)

	return fiber.Config{
		AppName:       "portfolio api",
		ServerHeader:  "your mother",
		Prefork:       prefork,
		CaseSensitive: caseSensitive,
		StrictRouting: strictRouting,
		ReadTimeout:   time.Second * time.Duration(readTimeout),
	}
}
