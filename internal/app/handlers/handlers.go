package handlers

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{Message: message})
}

func SendSuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}
