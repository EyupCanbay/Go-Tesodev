package repo

import (
	"context"
	"tesodev/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	_, err := r.Collection.InsertOne(ctx, product)
	return err
}
