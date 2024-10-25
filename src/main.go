package main

import (
	handler "golphin/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	api := app.Group("/api")
	api.POST("", handler.GenerateContent)
	app.Logger.Fatal(app.Start(":8080")) 
}