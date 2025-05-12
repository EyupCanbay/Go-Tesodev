package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"tesodev/models"
	"tesodev/repo"
	"tesodev/services"
	"tesodev/utils"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

const timeOut = 10 * time.Second

func CreateProduct(t *testing.T) models.Product {
	productCollection := testDB.Collection("products")

	productRepo := &repo.ProductRepository{
		Collection: productCollection,
	}

	productService := &services.ProductService{
		Repo: productRepo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	req := services.CreateProductRequest{
		Name:        utils.RandomString(10),
		Price:       float64(utils.RandomInt(10, 900)),
		Description: utils.RandomString(25),
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

	return retrieveProduct
}

func TestCreateProductService(t *testing.T) {
	CreateProduct(t)
}

func TestGetOneProductService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	product := CreateProduct(t)
	fmt.Print(product)

	productionCollection := testDB.Collection("products")

	productRepo := repo.ProductRepository{
		Collection: productionCollection,
	}

	productService := services.ProductService{
		Repo: &productRepo,
	}

	objId := product.Id.Hex()

	retrieveProduct, err := productService.GetOneId(ctx, objId)
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, retrieveProduct)
	require.Equal(t, product.Id, retrieveProduct.ProductId)
	require.Equal(t, product.Name, retrieveProduct.Name)
	require.Equal(t, product.Description, product.Description)
	require.Equal(t, product.Price, retrieveProduct.Price)
	require.WithinDuration(t, product.Created_at, retrieveProduct.Created_at, time.Second)

}

func TestPutProductService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	collectionProduct := testDB.Collection("products")

	productRepo := repo.ProductRepository{
		Collection: collectionProduct,
	}

	productSrevice := services.ProductService{
		Repo: &productRepo,
	}

	newProduct := CreateProduct(t)

	updatedProduct := services.CreateProductRequest{
		Name:        utils.RandomString(10),
		Price:       float64(utils.RandomPrice()),
		Description: newProduct.Description,
	}

	id := newProduct.Id.Hex()
	product, err := productSrevice.Update(ctx, id, updatedProduct)

	require.NoError(t, err)
	require.NotEmpty(t, product.ProductId)

	var retrieveProduct models.Product
	err = collectionProduct.FindOne(ctx, bson.M{"_id": newProduct.Id}).Decode(&retrieveProduct)
	require.NoError(t, err)
	require.NotEqual(t, newProduct.Name, retrieveProduct.Name)
	require.NotEqual(t, newProduct.Price, retrieveProduct.Price)
	require.Equal(t, retrieveProduct.Description, newProduct.Description)
	require.Equal(t, newProduct.Id, retrieveProduct.Id)
	require.WithinDuration(t, newProduct.Created_at, retrieveProduct.Created_at, timeOut)

}

func TestPatchProductService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	collectionProduct := testDB.Collection("products")

	productRepo := repo.ProductRepository{
		Collection: collectionProduct,
	}

	productService := services.ProductService{
		Repo: &productRepo,
	}

	// for name changed
	newProduct := CreateProduct(t)
	updatedProduct := services.CreateProductRequest{
		Name: utils.RandomString(15),
	}
	id := newProduct.Id.Hex()

	product, err := productService.Patch(ctx, id, updatedProduct)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	var result models.Product
	err = collectionProduct.FindOne(ctx, bson.M{"_id": product.ProductId}).Decode(&result)
	require.NotEqual(t, result.Name, newProduct.Name)
	require.Equal(t, result.Id, newProduct.Id)
	require.Equal(t, result.Description, newProduct.Description)
	require.Equal(t, newProduct.Price, result.Price)
	require.WithinDuration(t, newProduct.Created_at, result.Created_at, time.Second)

	// for description changed

	newProduct = CreateProduct(t)
	updatedProduct = services.CreateProductRequest{
		Description: utils.RandomString(15),
	}
	id = newProduct.Id.Hex()

	product, err = productService.Patch(ctx, id, updatedProduct)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	err = collectionProduct.FindOne(ctx, bson.M{"_id": product.ProductId}).Decode(&result)
	require.Equal(t, result.Name, newProduct.Name)
	require.Equal(t, result.Id, newProduct.Id)
	require.NotEqual(t, result.Description, newProduct.Description)
	require.Equal(t, newProduct.Price, result.Price)
	require.WithinDuration(t, newProduct.Created_at, result.Created_at, time.Second)

	// for price changed
	newProduct = CreateProduct(t)
	updatedProduct = services.CreateProductRequest{
		Price: float64(utils.RandomPrice()),
	}
	id = newProduct.Id.Hex()

	product, err = productService.Patch(ctx, id, updatedProduct)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	err = collectionProduct.FindOne(ctx, bson.M{"_id": product.ProductId}).Decode(&result)
	require.Equal(t, result.Name, newProduct.Name)
	require.Equal(t, result.Id, newProduct.Id)
	require.Equal(t, result.Description, newProduct.Description)
	require.NotEqual(t, newProduct.Price, result.Price)
	require.WithinDuration(t, newProduct.Created_at, result.Created_at, time.Second)
}

func TestGetProductService(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateProduct(t)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	collectionProduct := testDB.Collection("products")

	productRepo := repo.ProductRepository{
		Collection: collectionProduct,
	}

	productService := services.ProductService{
		Repo: &productRepo,
	}

	opt := services.CreateProductRequest{
		Limit: 10,
		Page:  1,
	}

	results, err := productService.Get(ctx, opt)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	for _, product := range results {
		require.NotEmpty(t, product)
	}
}

func TestDeleteProduct(t *testing.T) {
	newProduct := CreateProduct(t)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	collectionProduct := testDB.Collection("products")

	productRepo := repo.ProductRepository{
		Collection: collectionProduct,
	}

	productService := services.ProductService{
		Repo: &productRepo,
	}

	id := newProduct.Id.Hex()
	err := productService.Delete(ctx, id)
	require.NoError(t, err)

	var product models.Product
	err = collectionProduct.FindOne(ctx, bson.M{"_id": newProduct.Id}).Decode(&product)
	require.Error(t, err)
	require.Empty(t, product)

}
