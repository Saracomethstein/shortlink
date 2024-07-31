package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shortlink/internal/app/models"
	"shortlink/internal/app/services"
)

type IAuthHandler interface {
	Authorization(c echo.Context) error
	Register(c echo.Context) error
}

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Authorization(c echo.Context) error {
	var req models.LoginRequest
	var err error

	if req, err = models.BindLogin(c); err != nil {
		return SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	token, err := h.AuthService.Authorization(req.Username, req.Password)
	if err != nil || token == "" {
		return SendErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "true", "session_id": token})
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}

	err := h.AuthService.Registration(req.Username, req.Password)
	if err != nil {
		return SendErrorResponse(c, http.StatusInternalServerError, "Could not register user")
	}

	return SendSuccessResponse(c, http.StatusOK, "User registered successfully")
}
