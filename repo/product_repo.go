package repo

import (
	"context"
	"fmt"
	"tesodev/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) (*primitive.ObjectID, error) {

	now := time.Now()
	product.Created_at = now

	result, err := r.Collection.InsertOne(ctx, product)

	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("result is empty")
	}
	id, ok := result.InsertedID.(primitive.ObjectID) //type assertion
	if !ok {
		return nil, fmt.Errorf("Invalid id type")
	}

	return &id, err
}

func (r *ProductRepository) GetOneId(ctx context.Context, id string) (*models.Product, error) {
	ObcID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = r.Collection.FindOne(ctx, bson.M{"_id": ObcID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Get(ctx context.Context) ([]models.Product, error) {
	var products []models.Product

	result, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		var singleProduct models.Product
		if err = result.Decode(&singleProduct); err != nil {
			return nil, err
		}

		products = append(products, singleProduct)
	}

	return products, nil
}

func (r *ProductRepository) Update(ctx context.Context, id string, updateProduct bson.M) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateProduct})
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no document found with the given id")
	}
	return &models.Product{
		Id: objID,
	}, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("did not delete the product")
	}

	return nil
}

func (r *ProductRepository) Patch(ctx context.Context, id string, updatePrduct bson.M) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updatePrduct["updated_at"] = time.Now()
	filter := bson.M{"_id": objID}
	updateData := bson.M{"$set": updatePrduct}

	result, err := r.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no document found with the given id")
	}
	return &models.Product{
		Id: objID,
	}, nil
}
