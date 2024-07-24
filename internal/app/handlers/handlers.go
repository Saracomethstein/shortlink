package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shortlink/internal/app/dbConnection"
)

type Request struct {
	OriginalURL string `json:"url"`
}

type Response struct {
	ShortURL string `json:"short-url"`
}

func HandlerHi(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HandlerAddUrl(c echo.Context) error {
	var req Request
	var shortUrl string
	err := dbConnection.AddUrl(req.OriginalURL, shortUrl)

	if err != nil {
		return err
	}

	fmt.Println("Urls added in db.")
	return nil
}
