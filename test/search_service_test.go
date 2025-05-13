package test

import (
	"context"
	"tesodev/models"
	"tesodev/repo"
	"tesodev/services"
	"tesodev/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func DbFormater(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	productCollection := testDB.Collection("products")

	result, err := productCollection.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)
	require.NotEqual(t, result.DeletedCount, 0)
}

func CreateCertainProduct(t *testing.T, req services.CreateProductRequest) models.Product {

	productCollection := testDB.Collection("products")

	productRepo := &repo.ProductRepository{
		Collection: productCollection,
	}

	productService := &services.ProductService{
		Repo: productRepo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

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

func TestSearchService(t *testing.T) {

	DbFormater(t)

	productCollection := testDB.Collection("products")

	productRepo := &repo.ProductRepository{
		Collection: productCollection,
	}

	productService := &services.ProductService{
		Repo: productRepo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	// 1. search is name params running for exact name
	req := services.CreateProductRequest{
		Name:        "Laptopp",
		Price:       12000,
		Description: utils.RandomString(20),
	}

	newProduct1 := CreateCertainProduct(t, req)

	search1 := services.CreateProductSearchRequest{
		Name: "Laptopp",
	}

	searchResult1, err := productService.SearchProducts(ctx, search1)
	require.NoError(t, err)
	require.NotEmpty(t, searchResult1)
	require.Len(t, searchResult1, 1)
	require.Equal(t, newProduct1.Id, searchResult1[0].ProductId)
	require.Equal(t, newProduct1.Name, searchResult1[0].Name)
	require.Equal(t, newProduct1.Description, searchResult1[0].Description)
	require.Equal(t, newProduct1.Price, searchResult1[0].Price)
	require.WithinDuration(t, newProduct1.Created_at, searchResult1[0].Created_at, time.Second)

	// 2. search is min price params running correctly
	// use the same request with up search

	search2 := services.CreateProductSearchRequest{
		PriceMin: 19000000,
	}
	searchResult2, err := productService.SearchProducts(ctx, search2)
	require.Error(t, err)
	require.Empty(t, searchResult2)
	require.Len(t, searchResult2, 0)

	// 3. search are name parms and max price running.  Name params is for partial name (with regex search)

	search3 := services.CreateProductSearchRequest{
		Name:     "lapt",
		PriceMax: 1900000,
	}
	searchResult3, err := productService.SearchProducts(ctx, search3)
	require.NoError(t, err)
	require.NotEmpty(t, searchResult3)
	require.Len(t, searchResult3, 1)
	require.Equal(t, newProduct1.Id, searchResult3[0].ProductId)
	require.Equal(t, newProduct1.Name, searchResult3[0].Name)
	require.Equal(t, newProduct1.Description, searchResult3[0].Description)
	require.Equal(t, newProduct1.Price, searchResult3[0].Price)
	require.WithinDuration(t, newProduct1.Created_at, searchResult3[0].Created_at, time.Second)

	// 4. search are  max and min price search running
	search4 := services.CreateProductSearchRequest{
		PriceMax: 100,
		PriceMin: 1,
	}
	searchResult4, err := productService.SearchProducts(ctx, search4)
	require.Error(t, err)
	require.Empty(t, searchResult4)
	require.Len(t, searchResult4, 0)

	// 5. search is the empty search running
	search5 := services.CreateProductSearchRequest{}
	searchResult5, err := productService.SearchProducts(ctx, search5)
	require.NoError(t, err)
	require.NotEmpty(t, searchResult5)
	require.Len(t, searchResult5, 1)
	require.Equal(t, newProduct1.Id, searchResult5[0].ProductId)
	require.Equal(t, newProduct1.Name, searchResult5[0].Name)
	require.Equal(t, newProduct1.Description, searchResult5[0].Description)
	require.Equal(t, newProduct1.Price, searchResult5[0].Price)
	require.WithinDuration(t, newProduct1.Created_at, searchResult5[0].Created_at, time.Second)

	// 6. search is name params running for a lot of documants
	req = services.CreateProductRequest{
		Name:        "LaptopBramds",
		Price:       115000,
		Description: utils.RandomString(20),
	}

	newProduct2 := CreateCertainProduct(t, req)

	search6 := services.CreateProductSearchRequest{
		Name:     "Lapt",
		PriceMax: 1900000,
		PriceMin: 1300,
	}

	searchResult6, err := productService.SearchProducts(ctx, search6)
	require.NoError(t, err)
	require.NotEmpty(t, searchResult6)
	require.Len(t, searchResult6, 2)
	require.Equal(t, newProduct1.Id, searchResult6[0].ProductId)
	require.Equal(t, newProduct1.Name, searchResult6[0].Name)
	require.Equal(t, newProduct1.Description, searchResult6[0].Description)
	require.Equal(t, newProduct1.Price, searchResult6[0].Price)
	require.WithinDuration(t, newProduct1.Created_at, searchResult6[0].Created_at, time.Second)

	require.Equal(t, newProduct2.Id, searchResult6[1].ProductId)
	require.Equal(t, newProduct2.Name, searchResult6[1].Name)
	require.Equal(t, newProduct2.Description, searchResult6[1].Description)
	require.Equal(t, newProduct2.Price, searchResult6[1].Price)
	require.WithinDuration(t, newProduct2.Created_at, searchResult6[1].Created_at, time.Second)

	//7. search while tehre are a lot of prdocuts is running exact name search
	search7 := services.CreateProductSearchRequest{
		Name:     "Laptopp",
		PriceMax: 1900000,
	}
	searchResult7, err := productService.SearchProducts(ctx, search7)
	require.NoError(t, err)
	require.NotEmpty(t, searchResult7)
	require.Len(t, searchResult7, 1)
	require.Equal(t, newProduct1.Id, searchResult7[0].ProductId)
	require.Equal(t, newProduct1.Name, searchResult7[0].Name)
	require.Equal(t, newProduct1.Description, searchResult7[0].Description)
	require.Equal(t, newProduct1.Price, searchResult7[0].Price)
	require.WithinDuration(t, newProduct1.Created_at, searchResult7[0].Created_at, time.Second)

}
