package main

import (
	"github.com/labstack/echo/v4"
	"shortlink/internal/app/dbConnection"
	"shortlink/internal/app/handlers"
)

func main() {
	dbConnection.GetConnection()
	defer dbConnection.CloseConnection()

	e := echo.New()
	e.Static("/", "./website/static")
	e.GET("/hi", handlers.HandlerHi)
	e.GET("/shortlink", handlers.HandlerAddUrl)
	e.Logger.Fatal(e.Start(":8000"))
}
