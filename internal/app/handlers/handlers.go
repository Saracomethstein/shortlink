package handlers

import (
	"net/http"
	"net/url"
	"shortlink/internal/app/dbConnection"

	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}

func HandlerAddUrl(c echo.Context) error {
	var req URLRequest
	var shortID string

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL"})
	}

	err, answer := dbConnection.CheckURLExists(req.URL)

	if err != nil {
		return err
	}

	if answer == true {
		err, shortID = dbConnection.GetShortURL(req.URL)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}
	} else {
		shortID, err = shortid.Generate()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating short ID"})
		}

		err = dbConnection.AddUrl(req.URL, shortID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving URL"})
		}
	}
	response := URLResponse{ShortenedURL: "http://localhost:8000/" + shortID}
	return c.JSON(http.StatusOK, response)
}

func HandlerRedirect(c echo.Context) error {
	shortURL := c.Param("shortURL")

	_, originalURL := dbConnection.GetUrl(c, shortURL)

	return c.Redirect(http.StatusFound, originalURL)
}
