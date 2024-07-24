package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerHi(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
