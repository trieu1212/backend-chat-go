package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	dbName := os.Getenv("MONGO_DB")
	host := "localhost"
	port := os.Getenv("MONGO_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin&readPreference=primary&directConnection=true&ssl=false",
		username, password, host, port, dbName)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("trieu").Collection(collectionName)
}
