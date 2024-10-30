package handler

import (
	"golphin/src/api/dto/request"
	"golphin/src/api/dto/response"
	service "golphin/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GenerateContent é o manipulador para a geração de conteúdo baseado em um prompt.
// @Summary Gera conteúdo a partir de um prompt
// @Description Recebe um prompt e gera conteúdo correspondente utilizando um serviço.
// @Tags Content
// @Accept json
// @Produce json
// @Param payload body request.Gemini true "Payload contendo o prompt"
// @Success 200 {object} response.Gemini
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /generate [post]
func GenerateContent(c echo.Context) error {
	var payload request.Gemini
	if err := c.Bind(&payload); err != nil {
        return c.JSON(400, map[string]string{"error": err.Error()})
    }
	
	content, err := service.GenerateContent(payload.GetPrompt())
	if err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }

	response := response.NewGemini()
	response.SetContent(content)
	
    return c.JSON(http.StatusOK, response)
}
