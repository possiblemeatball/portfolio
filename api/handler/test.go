package handler

import (
	"github.com/chenmingyong0423/go-mongox"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/builder/update"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
	"possiblemeatball.com/api/v2/database"
	"possiblemeatball.com/api/v2/models"
)

func AddJohn(c *fiber.Ctx) error {
	mongoCollection := database.NewMongoCollection("names")
	collection := mongox.NewCollection[models.Name](mongoCollection)
	john := models.Name{
		FirstName: "John",
		LastName:  "Doe",
	}
	result, err := collection.Creator().InsertOne(context.Background(), &john)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(result)
}

func ChangeJohn(c *fiber.Ctx) error {
	mongoCollection := database.NewMongoCollection("names")
	collection := mongox.NewCollection[models.Name](mongoCollection)
	result, err := collection.Updater().Filter(query.Eq("first_name", "John")).Updates(update.Set("last_name", "Deere")).UpdateOne(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(result)
}
