package main

import (
	"shortlink/internal/app/dbConnection"
	"shortlink/internal/app/handlers"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	dbConnection.GetConnection()
	defer dbConnection.CloseConnection()

	e := echo.New()
	e.Static("/", "./website/static/auth")
	e.Static("/registration", "./website/static/regist")
	e.Static("/output", "./website/static/output")
	e.Static("/shorten", "./website/static/main")
	e.Static("/profile", "./website/static/profile")

	e.POST("/shorten", handlers.HandlerAddUrl)
	e.POST("/auth", handlers.HandlerAuth)
	e.POST("/registration", handlers.HandlerRegistration)
	e.GET("/redirect/:shortURL", handlers.HandlerRedirect)
	e.Logger.Fatal(e.Start(":8000"))
}
