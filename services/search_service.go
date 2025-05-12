package services

import (
	"context"
	"fmt"
	"tesodev/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateProductSearchRequest struct {
	Name     string
	PriceMin float64
	PriceMax float64
	Sort     string // "asc" or "desc"
	Limit    int
	Page     int
}

func (s *ProductService) SearchProducts(ctx context.Context, params CreateProductSearchRequest) ([]CreateProductResponse, error) {
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

	skip := (params.Page - 1) * params.Limit
	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(params.Limit))

	if params.Sort == "asc" {
		findOptions.SetSort(bson.D{{Key: "price", Value: 1}})
	} else if params.Sort == "desc" {
		findOptions.SetSort(bson.D{{Key: "price", Value: -1}})
	}

	if params.Name != "" {
		filter["name"] = params.Name
	}
	products, err := s.Repo.SearchProducts(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	buffer := len(products)
	if buffer == 0 && params.Name != "" {
		filter["name"] = bson.M{"$regex": params.Name, "$options": "i"}
		products, err := s.Repo.SearchProducts(ctx, filter, findOptions)
		if err != nil {
			return nil, err
		}
		result, _ := RevertToSlice(products)
		if len(result) == 0 {
			return nil, fmt.Errorf("Do not found that the name")
		}
	}
	return RevertToSlice(products)
}

func RevertToSlice(products []models.Product) ([]CreateProductResponse, error) {
	var result []CreateProductResponse
	for _, p := range products {
		result = append(result, CreateProductResponse{
			ProductId:   p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Created_at:  p.Created_at,
		})
	}
	return result, nil
}
