package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"shortlink/internal/app/models"
	"shortlink/internal/app/repositories"
)

type ProfileService struct {
	profileRepo repositories.ProfileRepository
}

func NewProfileService(profileRepo repositories.ProfileRepository) *ProfileService {
	return &ProfileService{profileRepo: profileRepo}
}

func (p *ProfileService) ProfileHistory(sission_id string) ([]models.Link, error) {
	var urls []models.Link
	var err error

	login, err := p.profileRepo.GetLoginFromLog(sission_id)
	if err != nil {
		return []models.Link{}, err
	}

	urls, err = p.profileRepo.GetUserHistory(login)
	if err != nil {
		return []models.Link{}, err
	}

	return urls, nil
}

func (p *ProfileService) GetSessionID(c echo.Context) (string, error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return "", err
	}
	session_id := cookie.Value
	fmt.Println(session_id)
	return session_id, nil
}
