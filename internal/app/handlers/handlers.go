package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"net/http"
	"net/url"
	"shortlink/internal/app/dbConnection"
)

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}

func HandlerAddUrl(c echo.Context) error {
	var req URLRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL"})
	}

	shortID, err := shortid.Generate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating short ID"})
	}
	err = dbConnection.AddUrl(req.URL, shortID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving URL"})
	}
	response := URLResponse{ShortenedURL: "http://localhost:8000/" + shortID}
	return c.JSON(http.StatusOK, response)
}

func HandlerRedirect(c echo.Context) error {
	shortID := c.Param("shortID")
	var originalURL string

	err := dbConnection.GetUrl(c, shortID, &originalURL)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, originalURL)
}
