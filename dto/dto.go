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
