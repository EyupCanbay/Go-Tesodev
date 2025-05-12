package test

import (
	"context"
	"testing"
	"time"

	"tesodev/models"
	"tesodev/repo"
	"tesodev/services"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateProductIntegration(t *testing.T) {
	productCollection := testDB.Collection("products")

	productRepo := &repo.ProductRepository{
		Collection: productCollection,
	}

	productService := &services.ProductService{
		Repo: productRepo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := services.CreateProductRequest{
		Name:        "Test Real Product",
		Price:       99.99,
		Description: "This is a real integration test product",
	}

	result, err := productService.Create(ctx, req)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.ProductId)

	var retrieveProduct models.Product
	err = productCollection.FindOne(ctx, bson.M{"_id": result.ProductId}).Decode(&retrieveProduct)
	require.NoError(t, err)
	require.Equal(t, req.Name, retrieveProduct.Name)
	require.Equal(t, req.Price, retrieveProduct.Price)
	require.Equal(t, retrieveProduct.Description, req.Description)

}
