package dto

import (
	"github.com/labstack/echo/v4"
)

func ErrorHandling(c echo.Context, statusCode int, err *echo.Map) error {
	message := false

	errorResponse := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       err,
	}

	return c.JSON(statusCode, errorResponse)
}

func SuccessHandling(c echo.Context, statusCode int, data *echo.Map) error {
	message := true

	successResponse := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	return c.JSON(statusCode, successResponse)
}
