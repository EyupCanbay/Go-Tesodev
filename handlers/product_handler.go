package handlers

import (
	"fmt"
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
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}

	if req.Name == "" || req.Description == "" {
		fmt.Println(req.Description, req.Name)
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "must be all feild required"})
	}

	if req.Price < 0 {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "sould be positife price"})
	}

	product := services.CreateProductRequest{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	result, err := h.Services.Create(c.Request().Context(), product)
	if err != nil {
		dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": result.ProductId})
}

func (h *ProductHandler) GetAProductId(c echo.Context) error {
	id := c.Param("product_id")

	if id == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "Id paramater must be feild"})
	}

	product, err := h.Services.GetOneId(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusNotFound, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": product})
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	products, err := h.Services.Get(c.Request().Context())
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": products})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("product_id")

	if id == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "Id paramater must be feild"})
	}
	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}

	if req.Description == "" || req.Name == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "must be all feild required"})
	}
	if req.Price < 0 {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "sould be positife price"})
	}

	serviceProduct := services.CreateProductRequest{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	result, err := h.Services.Update(c.Request().Context(), id, serviceProduct)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": result.ProductId})

}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("product_id")
	if id == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "Id paramater must be feild"})
	}

	err := h.Services.Delete(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": "Succesfuly delete the product"})
}

func (h *ProductHandler) UpdateSingleFeild(c echo.Context) error {
	id := c.Param("product_id")
	if id == "" {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": "Id paramater must be feild"})
	}
	var req dto.ProductRequest

	if err := c.Bind(&req); err != nil {
		return dto.ErrorHandling(c, http.StatusBadRequest, &echo.Map{"data": err.Error()})
	}

	serviceReq := services.CreateProductRequest{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	result, err := h.Services.Patch(c.Request().Context(), id, serviceReq)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"data": result.ProductId})
}
