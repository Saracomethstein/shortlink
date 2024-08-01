package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"shortlink/internal/app/handlers"
	"shortlink/internal/app/repositories"
	"shortlink/internal/app/services"
)

func main() {
	db := repositories.SetupDB()
	serviceContainer := services.NewServiceContainer(db)

	e := echo.New()

	e.Static("/", "./website/static/auth")
	e.Static("/registration", "./website/static/regist")
	e.Static("/output", "./website/static/output")
	e.Static("/shorten", "./website/static/main")
	e.Static("/profile", "./website/static/profile")

	authHandler := handlers.NewAuthHandler(serviceContainer.AuthService)
	linkHandler := handlers.NewLinkHandler(serviceContainer.LinkService)
	profileHandler := handlers.NewProfileHandler(serviceContainer.ProfileService)

	e.POST("/auth", authHandler.Authorization)
	e.POST("/registration", authHandler.Register)
	e.POST("/shorten", linkHandler.CreateShortLink)
	e.GET("/redirect/:shortCode", linkHandler.Redirect)
	e.GET("/profile", profileHandler.GetProfileData)

	e.Logger.Fatal(e.Start(":8000"))
}
