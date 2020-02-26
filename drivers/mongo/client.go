package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var (
	instance *mongo.Client
)

// ConnectSingleton returns reusable mongo client
func ConnectSingleton() *mongo.Client {
	if instance == nil {
		instance = getMongoClient()
	}
	return instance
}

func getMongoClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	mongoURL := os.Getenv("MONGO_LOCAL_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
