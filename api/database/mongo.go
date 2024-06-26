package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"possiblemeatball.com/api/v2/config"
)

func NewMongoCollection(collectionName string) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), config.EnvToMongoClientOptions())
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	collection := config.EnvToMongoCollection(client, collectionName)
	return collection
}
