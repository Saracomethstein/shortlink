package handlers

import (
	"fmt"
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

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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

	fmt.Println(shortID)

	response := URLResponse{ShortenedURL: "http://localhost:8000/redirect/" + shortID}
	return c.JSON(http.StatusOK, response)
}

func HandlerRedirect(c echo.Context) error {
	shortURL := c.Param("shortURL")

	_, originalURL := dbConnection.GetUrl(c, shortURL)

	return c.Redirect(http.StatusFound, originalURL)
}

func HandlerAuth(c echo.Context) error {
	var user User
	var answer bool
	var err error

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err, answer = dbConnection.CheckUserIdDB(user.Login, user.Password)

	if err != nil {
		return err
	}

	if answer == true {
		return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid login or password"})
	}
}

func HandlerRegistration(c echo.Context) error {
	var user User
	var err error

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err = dbConnection.AddNewUserInDB(user.Login, user.Password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Add user successful"})
}
