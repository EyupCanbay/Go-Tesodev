package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var timeOut = 5 * time.Second

func ConnectDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DatabaseEnv()))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongoDB")
	return client
}

func GetCollection(client *mongo.Client, col string) *mongo.Collection {
	collection := client.Database(os.Getenv("DATABASE")).Collection(col)
	return collection
}
