package handler

import (
	"golphin/src/api/dto/response"
	service "golphin/src/services"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GenerateContent(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
        return c.JSON(400, map[string]string{"error": err.Error()})
    }

    prompt, ok := payload["prompt"].(string)
	if !ok {
    	return c.JSON(400, map[string]string{"error": "Invalid payload format"})
    }
	content, err := service.GenerateContent(prompt)
	if err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }

	response := response.NewGemini()
	response.SetContent(content)
	
    return c.JSON(http.StatusOK, response)
}
