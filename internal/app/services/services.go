package services

import (
	"database/sql"
	"shortlink/internal/app/repositories"
)

type ServiceContainer struct {
	AuthService    *AuthService
	LinkService    *LinkService
	ProfileService *ProfileService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	userRepo := repositories.NewUserRepository(db)
	linkRepo := repositories.NewLinkRepository(db)
	profileRepo := repositories.NewProfileRepository(db)

	return &ServiceContainer{
		AuthService:    NewAuthService(*userRepo),
		LinkService:    NewLinkService(*linkRepo),
		ProfileService: NewProfileService(*profileRepo),
	}
}
