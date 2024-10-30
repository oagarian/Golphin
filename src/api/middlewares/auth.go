package middleware

import (
	"golphin/src/utils/environment"
	"github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func AuthJWT() echo.MiddlewareFunc {
    en, err := environment.Get()
    if err != nil {
        panic(err)
    }
    return echojwt.WithConfig(echojwt.Config{
        SigningKey:    []byte(en.JWTSecret),
        SigningMethod: "HS256",
    })
}
