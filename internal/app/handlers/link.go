package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shortlink/internal/app/models"
	"shortlink/internal/app/services"
)

type ILinkHandler interface {
	CreateShortLink(c echo.Context) error
	Redirect(c echo.Context) error
}

type LinkHandler struct {
	LinkService *services.LinkService
}

func NewLinkHandler(linkService *services.LinkService) *LinkHandler {
	return &LinkHandler{LinkService: linkService}
}

func (h *LinkHandler) CreateShortLink(c echo.Context) error {
	var req models.CreateShortLinkRequest
	var session_id string
	var err error

	if req, err = models.BindLink(c); err != nil {
		return SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	if session_id, err = h.LinkService.GetSessionID(c); err != nil {
		return err
	}
	shortCode, err := h.LinkService.ShortUrl(session_id, req.OriginalURL)
	if err != nil {
		return SendErrorResponse(c, http.StatusInternalServerError, "Could not create short link")
	}

	response := models.CreateShortLinkResponse{ShortCode: "http://localhost:8000/redirect/" + shortCode}
	return c.JSON(http.StatusOK, response)
}

func (h *LinkHandler) Redirect(c echo.Context) error {
	shortCode := c.Param("shortCode")

	originalURL, err := h.LinkService.Redirect(shortCode)
	if err != nil {
		return SendErrorResponse(c, http.StatusNotFound, "Short link not found")
	}

	return c.Redirect(http.StatusMovedPermanently, originalURL)
}
