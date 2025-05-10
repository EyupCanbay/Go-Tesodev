package services

import (
	"context"
	"tesodev/dto"
	"tesodev/models"
	"tesodev/repo"
	"time"
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
