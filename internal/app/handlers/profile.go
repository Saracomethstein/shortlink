package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"shortlink/internal/app/models"
	"shortlink/internal/app/services"
)

type IProfileHandler interface {
	GetProfileData(c echo.Context) error
}

type ProfileHandler struct {
	ProfileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{ProfileService: profileService}
}

func (h *ProfileHandler) GetProfileData(c echo.Context) error {
	var urlHistory []models.Link
	var session_id string
	var login string
	var err error

	if session_id, err = h.ProfileService.GetSessionID(c); err != nil {
		return err
	}

	urlHistory, err = h.ProfileService.ProfileHistory(session_id)
	if err != nil {
		return err
	}

	if login, err = h.ProfileService.GetUsername(session_id); err != nil {
		return err
	}

	domainCount := make(map[string]int)
	for _, link := range urlHistory {
		parsedURL, err := url.Parse(link.OriginalLink)
		if err != nil {
			continue
		}
		domain := parsedURL.Host
		domainCount[domain]++
	}

	response := map[string]interface{}{
		"username":   login,
		"urlHistory": urlHistory,
		"domains":    domainCount,
	}

	return c.JSON(http.StatusOK, response)
}
