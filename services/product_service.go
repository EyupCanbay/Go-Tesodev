package services

import (
	"context"
	"tesodev/dto"
	"tesodev/models"
	"tesodev/repo"
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
