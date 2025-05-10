package dto

import (
	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type ServiceProduct struct {
	Name        string
	Price       float64
	Description string
}

type Response struct {
	StatusCode int       `json:"stauscode"`
	Message    bool      `json:"message"`
	Data       *echo.Map `json:"data"`
}

type ProductSearchParams struct {
	Name     string
	Exact    bool
	PriceMin float64
	PriceMax float64
	Sort     string // "asc" or "desc"
	Limit    int
	Page     int
}
