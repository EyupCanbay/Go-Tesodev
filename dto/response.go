package dto

import "github.com/labstack/echo/v4"

func ErrorHandling(c echo.Context, statusCode int, err string) error {
	message := false

	errorResponse := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	return c.JSON(statusCode, errorResponse)
}
