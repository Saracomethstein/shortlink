package services

import (
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"net/url"
	"shortlink/internal/app/repositories"
)

type ILinkService interface {
	ShortUrl(c echo.Context) error
	Redirect(c echo.Context) error
}

type LinkService struct {
	linkRepo repositories.LinkRepository
}

func NewLinkService(linkRepo repositories.LinkRepository) *LinkService {
	return &LinkService{linkRepo: linkRepo}
}

func (s *LinkService) ShortUrl(originalLink string) (string, error) {
	var shortLink string

	if _, err := url.ParseRequestURI(originalLink); err != nil {
		return "", err
	}

	answer, err := s.linkRepo.CheckLinkExistByOriginal(originalLink)

	if err != nil {
		return "", err
	}

	if answer == true {
		shortLink, err = s.linkRepo.GetShortLink(originalLink)

		if err != nil {
			return "", err
		}
	} else {
		shortLink, err = shortid.Generate()

		if err != nil {
			return "", err
		}

		err = s.linkRepo.CreateShortLink(originalLink, shortLink)
		if err != nil {
			return "", err
		}
	}
	return shortLink, nil
}

func (s *LinkService) Redirect(shortLink string) (string, error) {
	originalURL, err := s.linkRepo.GetOriginalLink(shortLink)

	if err != nil {
		return "", err
	}

	return originalURL, nil
}
