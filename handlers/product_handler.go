package handlers

import (
	"net/http"
	"tesodev/dto"
	"tesodev/services"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Services *services.ProductService
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req dto.ProductRequest

	if err := c.Bind(&req); err != nil {
		dto.ErrorHandling(c, http.StatusBadRequest, err.Error())
	}

	serviceProduct := &dto.ServiceProduct{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	responseProduct, err := h.Services.Create(c.Request().Context(), serviceProduct)
	if err != nil {
		dto.ErrorHandling(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, responseProduct)
}
