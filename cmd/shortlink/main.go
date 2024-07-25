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
	e.Static("/", "./website/static")
	e.POST("/shorten", handlers.HandlerAddUrl)
	e.GET("/shortID", handlers.HandlerRedirect)
	e.Logger.Fatal(e.Start(":8000"))
}