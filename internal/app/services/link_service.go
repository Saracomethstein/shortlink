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

func (s *LinkService) ShortUrl(sesion_id, originalLink string) (string, error) {
	var shortLink string
	var login string

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

		login, err = s.linkRepo.GetLoginFromLog(sesion_id)
		if err != nil {
			return "", err
		}

		err = s.linkRepo.CreateShortLink(login, shortLink, originalLink)
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

func (s *LinkService) GetSessionID(c echo.Context) (string, error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return "", err
	}
	session_id := cookie.Value
	return session_id, nil
}
