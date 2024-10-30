package api

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RichContext struct {
	echo.Context
	UserID uuid.UUID
}

func Bind(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
        ctx := &RichContext{Context: c}
        return next(ctx)
    }
}