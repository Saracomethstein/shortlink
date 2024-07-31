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

func (h *ProfileHandler) CreateHistory(c echo.Context) error {
	var urlHistory []models.Link
	var session_id string
	var err error

	if session_id, err = h.ProfileService.GetSessionID(c); err != nil {
		return err
	}

	urlHistory, err = h.ProfileService.ProfileHistory(session_id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, urlHistory)
}
