package handlers

import (
	"context"
	"net/http"
	"strconv"
	"tesodev/dto"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) Search(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := dto.ProductSearchParams{
		Name:  c.QueryParam("name"),
		Exact: c.QueryParam("exact") == "true",
		Sort:  c.QueryParam("sort"),
		Page:  1,
		Limit: 10,
	}

	if p := c.QueryParam("page"); p != "" {
		if pageInt, err := strconv.Atoi(p); err == nil && pageInt > 0 {
			params.Page = pageInt
		}
	}
	if l := c.QueryParam("limit"); l != "" {
		if limitInt, err := strconv.Atoi(l); err == nil && limitInt > 0 {
			params.Limit = limitInt
		}
	}
	if min := c.QueryParam("price_min"); min != "" {
		params.PriceMin, _ = strconv.ParseFloat(min, 64)
	}
	if max := c.QueryParam("price_max"); max != "" {
		params.PriceMax, _ = strconv.ParseFloat(max, 64)
	}

	products, err := h.Services.SearchProducts(ctx, params)
	if err != nil {
		return dto.ErrorHandling(c, http.StatusInternalServerError, &echo.Map{"data": err.Error()})
	}

	return dto.SuccessHandling(c, http.StatusOK, &echo.Map{"products": products})
}
