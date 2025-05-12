package services

import (
	"context"
	"fmt"
	"tesodev/models"
	"tesodev/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductService struct {
	Repo *repo.ProductRepository
}

type Product struct {
	Id          primitive.ObjectID
	Name        string
	Price       float64
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}

type CreateProductRequest struct {
	Name        string
	Price       float64
	Description string
	Limit       int
	Page        int
}

type CreateProductResponse struct {
	ProductId   primitive.ObjectID
	Name        string
	Price       float64
	Description string
	Created_at  time.Time
}

func (s *ProductService) Create(ctx context.Context, req CreateProductRequest) (*CreateProductResponse, error) {
	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}
	product.Created_at = time.Now()
	product.Updated_at = time.Now()

	id, err := s.Repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		ProductId: *id,
	}, nil
}

func (s *ProductService) GetOneId(ctx context.Context, id string) (*CreateProductResponse, error) {
	product, err := s.Repo.GetOneId(ctx, id)
	if err != nil || product == nil {
		return nil, fmt.Errorf("product not found or internal error")
	}
	return &CreateProductResponse{
		ProductId:   product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Created_at:  product.Created_at,
	}, nil
}

func (s *ProductService) Get(ctx context.Context, params CreateProductRequest) ([]CreateProductResponse, error) {

	skip := (params.Page - 1) * params.Limit

	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(params.Limit))

	products, err := s.Repo.Get(ctx, findOptions)
	if err != nil {
		return nil, err
	}

	var responses []CreateProductResponse
	for _, p := range products {
		responses = append(responses, CreateProductResponse{
			ProductId:   p.Id,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Created_at:  p.Created_at,
		})
	}

	return responses, nil
}

func (s *ProductService) Update(ctx context.Context, id string, req CreateProductRequest) (*CreateProductResponse, error) {

	var product = bson.M{}

	product["name"] = req.Name
	product["price"] = req.Price
	product["description"] = req.Description
	product["updated_at"] = time.Now()

	productData, err := s.Repo.Update(ctx, id, product)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		ProductId: productData.Id,
	}, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {

	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) Patch(ctx context.Context, id string, req CreateProductRequest) (*CreateProductResponse, error) {

	if req.Price < 0 {
		return nil, fmt.Errorf("Price do not  be negatife number")
	}

	updateProduct := bson.M{}
	if req.Name != "" {
		updateProduct["name"] = req.Name
	}
	if req.Price > 0 {
		updateProduct["price"] = req.Price
	}
	if req.Description != "" {
		updateProduct["description"] = req.Description
	}

	if len(updateProduct) == 0 {
		return nil, fmt.Errorf("No valid fields to updates")
	}

	product, err := s.Repo.Patch(ctx, id, updateProduct)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		ProductId: product.Id,
	}, nil
}
