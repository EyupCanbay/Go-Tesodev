package repo

import (
	"context"
	"tesodev/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	_, err := r.Collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) GetSingle(ctx context.Context, id string) (*models.Product, error) {
	ObcID, _ := primitive.ObjectIDFromHex(id)
	var product models.Product
	err := r.Collection.FindOne(ctx, bson.M{"_id": ObcID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
