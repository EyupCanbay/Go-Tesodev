package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Price       float64            `json:"price" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}

func CreateIndex(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mod := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: "text"},
		},
		Options: options.Index().SetName("TextIndex"),
	}

	c, err := collection.Indexes().CreateOne(ctx, mod)
	fmt.Println(c)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
