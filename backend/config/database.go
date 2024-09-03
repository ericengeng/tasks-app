package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to mongoDB:", err)
	}
	log.Println("Connected to MongoDB")
	return client
}
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("tasksDB").Collection(collectionName)
}
