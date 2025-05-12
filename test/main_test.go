package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	testClient *mongo.Client
	testDB     *mongo.Database
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := godotenv.Load("../.env")

	connectionString := os.Getenv("DB_URI")
	fmt.Println("fghgjjhjk")
	fmt.Println(connectionString)
	testClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	testDB = testClient.Database("tesodev_product")

	code := m.Run()

	if err := testClient.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	os.Exit(code)
}
