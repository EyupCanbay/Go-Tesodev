package services

import (
	"context"
	"errors"
	"tesodev/dto"
	"tesodev/models"
	"tesodev/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductService struct {
	Repo *repo.ProductRepository
}

func (s *ProductService) Create(ctx context.Context, req *dto.ServiceProduct) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}
	product.Created_at = time.Now()
	product.Updated_at = time.Now()

	if err := s.Repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ProductService) GetSingle(ctx context.Context, id string) (*models.Product, error) {
	product, err := s.Repo.GetSingle(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Update(ctx context.Context, id string, req *dto.ServiceProduct) error {

	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	product.Updated_at = time.Now()

	err := s.Repo.Update(ctx, id, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {

	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) Patch(ctx context.Context, id string, req *dto.ProductRequest) error {

	if req.Price < 0 {
		return errors.New("Price do not  be negatife number")
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
		return errors.New("No valid fields to updates")
	}

	return s.Repo.Patch(ctx, id, updateProduct)
}
