package services

import (
	"database/sql"
	"shortlink/internal/app/repositories"
)

type ServiceContainer struct {
	AuthService *AuthService
	LinkService *LinkService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	userRepo := repositories.NewUserRepository(db)
	linkRepo := repositories.NewLinkRepository(db)

	return &ServiceContainer{
		AuthService: NewAuthService(*userRepo),
		LinkService: NewLinkService(*linkRepo),
	}
}
