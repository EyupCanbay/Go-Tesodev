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

func (s *ProductService) Create(ctx context.Context, req *dto.ServiceProduct) (*dto.SuccessResponse, error) {
	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	if err := s.Repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Id:          product.Id.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}
