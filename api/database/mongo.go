package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

func NewMongoCollection() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "test",
		Password:      "test",
	}))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	collection := client.Database("test").Collection("test")
	return collection
}
