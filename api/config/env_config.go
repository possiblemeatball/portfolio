package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func EnvToMongoClientOptions() *options.ClientOptions {
	host := os.Getenv("MONGO_HOST")
	port, _ := strconv.ParseInt(os.Getenv("MONGO_PORT"), 10, 64)
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
	return options.Client().ApplyURI(uri)
}

func EnvToMongoCollection(client *mongo.Client, collection string) *mongo.Collection {
	database := os.Getenv("MONGO_DATABASE")
	return client.Database(database).Collection(collection)
}
