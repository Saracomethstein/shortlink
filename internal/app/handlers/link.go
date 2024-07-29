package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

type CreateShortLinkRequest struct {
	OriginalURL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ShortCode string `json:"shortenedUrl"`
}

func (h *LinkHandler) CreateShortLink(c echo.Context) error {
	var req CreateShortLinkRequest
	if err := c.Bind(&req); err != nil {
		return SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	shortCode, err := h.LinkService.ShortUrl(req.OriginalURL)
	if err != nil {
		return SendErrorResponse(c, http.StatusInternalServerError, "Could not create short link")
	}

	response := CreateShortLinkResponse{ShortCode: "http://localhost:8000/redirect/" + shortCode}
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
