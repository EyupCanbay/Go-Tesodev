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
		dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}

	serviceProduct := &dto.ServiceProduct{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	_, err := h.Services.Create(c.Request().Context(), serviceProduct)
	if err != nil {
		dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": "creted"})
}

func (h *ProductHandler) GetAProduct(c echo.Context) error {
	id := c.Param("product_id")

	if id == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "Id paramater must be feild"})
	}

	product, err := h.Services.GetSingle(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusNotFound, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": product})
}

func (h ProductHandler) GetAllProduct(c echo.Context) error {
	products, err := h.Services.GetAll(c.Request().Context())
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": products})
}
