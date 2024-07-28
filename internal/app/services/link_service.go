package services

import (
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"net/http"
	"net/url"
	"shortlink/internal/app/repositories"
)

type ILinkService interface {
	ShortUrl(c echo.Context) error
	Redirect(c echo.Context) error
}

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}

type LinkService struct {
	linkRepo repositories.LinkRepository
}

func NewLinkService(linkRepo repositories.LinkRepository) *LinkService {
	return &LinkService{linkRepo: linkRepo}
}

func (s *LinkService) ShortUrl(c echo.Context) error {
	var req URLRequest
	var shortLink string

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL"})
	}

	answer, err := s.linkRepo.CheckLinkExistByOriginal(req.URL)

	if err != nil {
		return err
	}

	if answer == true {
		shortLink, err = s.linkRepo.GetShortLink(req.URL)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}
	} else {
		shortLink, err = shortid.Generate()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating short ID"})
		}

		err = s.linkRepo.CreateShortLink(req.URL, shortLink)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving URL"})
		}
	}

	response := URLResponse{ShortenedURL: "http://localhost:8000/redirect/" + shortLink}
	return c.JSON(http.StatusOK, response)
}

func (s *LinkService) Redirect(c echo.Context) error {
	var link repositories.Link
	var err error

	link.ShortCode = c.Param("shortURL")

	link.OriginalURL, err = s.linkRepo.GetOriginalLink(link.ShortCode)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, link.OriginalURL)
}
