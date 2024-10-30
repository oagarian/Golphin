package main

import (
	"golphin/src/api"
	handler "golphin/src/api/handlers"
	middleware "golphin/src/api/middlewares"
	"github.com/labstack/echo/v4"
	_ "golphin/src/docs"
	"github.com/swaggo/echo-swagger"
)

// @title API Documentação
// @version 1.0
// @description Esta é a documentação da API do meu projeto.
// @host localhost:8080
// @BasePath /
func main() {
	app := echo.New()
	app.Use(api.Bind)
	api := app.Group("/api")
	api.Use(middleware.AuthJWT())
	api.POST("", handler.GenerateContent)
	auth := app.Group("/auth")
	auth.POST("/register", handler.AuthHandler) 
	auth.POST("/login", handler.LoginHandler)
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.Logger.Fatal(app.Start(":8080")) 
}