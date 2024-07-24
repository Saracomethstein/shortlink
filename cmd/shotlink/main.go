package main

import (
	"github.com/labstack/echo/v4"
	"shortlink/internal/app/handlers"
)

func main() {
	e := echo.New()

	e.Static("/", "./website/static")

	e.GET("/hi", handlers.HandlerHi)
	e.Logger.Fatal(e.Start(":8000"))
}
