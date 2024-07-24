package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"shortlink/internal/app/dbConnection"
	"shortlink/internal/app/handlers"
)

func main() {
	dbConnection.GetConnection()
	defer dbConnection.CloseConnection()

	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.Static("/", "./website/static")
	e.POST("/shorten", handlers.HandlerAddUrl)
	e.GET("/shortID", handlers.HandlerRedirect)
	e.Logger.Fatal(e.Start(":8000"))
}
