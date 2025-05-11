package services

import (
	"context"
	"fmt"
	"tesodev/dto"
	"tesodev/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *ProductService) SearchProducts(ctx context.Context, params dto.ProductSearchParams) ([]models.Product, error) {
	filter := bson.M{}

	if params.PriceMin > 0 || params.PriceMax > 0 {
		priceFilter := bson.M{}
		if params.PriceMin > 0 {
			priceFilter["$gte"] = params.PriceMin
		}
		if params.PriceMax > 0 {
			priceFilter["$lte"] = params.PriceMax
		}
		filter["price"] = priceFilter
	}

	if params.Name != "" {
		filter["name"] = params.Name
	}

	skip := (params.Page - 1) * params.Limit
	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(params.Limit))

	if params.Sort == "asc" {
		findOptions.SetSort(bson.D{{Key: "price", Value: 1}})
	} else if params.Sort == "desc" {
		findOptions.SetSort(bson.D{{Key: "price", Value: -1}})
	}

	fmt.Println("filter", filter)
	fmt.Println("find", findOptions)
	products, _ := s.Repo.SearchProducts(ctx, filter, findOptions)
	fmt.Println(products)
	buffer := len(products)
	if buffer == 0 {

		filter["name"] = bson.M{"$regex": params.Name, "$options": "i"}

		return s.Repo.SearchProducts(ctx, filter, findOptions)
	}
	return products, nil
}
