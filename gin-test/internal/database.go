package internal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URL        string
	DBName     string
	Collection string
}

var MongoClient *mongo.Client
var MongoCollection *mongo.Collection

func InitDatabase(config MongoConfig) {
	var err error
	clientOptions := options.Client().ApplyURI(config.URL)

	MongoClient, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = MongoClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	MongoCollection = MongoClient.Database(config.DBName).Collection(config.Collection)
	log.Println("MongoDB connected successfully!")
}
