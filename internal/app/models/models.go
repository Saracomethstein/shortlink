package models

import (
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type Link struct {
	ShortLink    string `json:"shortenedUrl"`
	OriginalLink string `json:"url"`
}

type CreateShortLinkRequest struct {
	OriginalURL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ShortCode string `json:"shortenedUrl"`
}

func BindLogin(c echo.Context) (LoginRequest, error) {
	var user LoginRequest

	if err := c.Bind(&user); err != nil {
		return LoginRequest{}, err
	}
	return user, nil
}

func BindLink(c echo.Context) (CreateShortLinkRequest, error) {
	var link CreateShortLinkRequest

	if err := c.Bind(&link); err != nil {
		return CreateShortLinkRequest{}, err
	}
	return link, nil
}
