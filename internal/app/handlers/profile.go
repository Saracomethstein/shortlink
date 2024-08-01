package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shortlink/internal/app/models"
	"shortlink/internal/app/services"
)

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

	response := map[string]interface{}{
		"username":   login,
		"urlHistory": urlHistory,
	}

	return c.JSON(http.StatusOK, response)
}
