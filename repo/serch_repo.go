package repo

import (
	"context"
	"tesodev/models"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *ProductRepository) SearchProducts(ctx context.Context, filter interface{}, opts *options.FindOptions) ([]models.Product, error) {

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var singleProduct models.Product
		if err = cursor.Decode(&singleProduct); err != nil {
			return nil, err
		}

		products = append(products, singleProduct)
	}

	return products, nil
}
