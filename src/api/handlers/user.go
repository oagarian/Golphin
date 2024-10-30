package handler

import (
	"fmt"
	"golphin/src/domain/user"
	"golphin/src/repository"
	"golphin/src/utils/encrypt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// AuthHandler é o manipulador para a criação de novos usuários.
// @Summary Cria um novo usuário
// @Description Cria um novo usuário no sistema com um nome de usuário, email e senha.
// @Tags Auth
// @Accept json
// @Produce json
// @Param username formData string true "Nome de usuário"
// @Param email formData string true "Email do usuário"
// @Param password formData string true "Senha do usuário"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func AuthHandler( c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
    password := c.FormValue("password")
	if username == "" || email == "" || password == "" {
        return c.JSON(400, map[string]string{"error": "Missing required fields"})
    }
	if len(password) < 8 || len(password) > 32 {
		return c.JSON(400, map[string]string{"error": "Password must be between 8 and 32 characters"})
    }

	ecryptedPassword, err := encrypt.HashPassword(password)
	if err != nil {
        return c.JSON(500, map[string]string{"error": "Internal Server Error"})
    }
	userBuilder := user.NewBuilder()
	userBuilder.
		SetID(uuid.New()).
		SetUsername(username).
		SetEmail(email).
		SetPassword(ecryptedPassword)
	user, err := userBuilder.Build()
	if err != nil {
        return c.JSON(400, map[string]string{"error": err.Error()})
    }
    err = repository.Auth(user)
	if err != nil {
		fmt.Println(err)
        return c.JSON(500, map[string]string{"error": "Internal Server Error"})
    }
    return c.NoContent(201)
}

// LoginHandler é o manipulador para a autenticação de usuários.
// @Summary Realiza o login do usuário
// @Description Autentica um usuário com nome de usuário ou email e senha.
// @Tags Auth
// @Accept json
// @Produce json
// @Param auth formData string true "Nome de usuário ou email"
// @Param password formData string true "Senha do usuário"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func LoginHandler(c echo.Context) error {
    auth := c.FormValue("auth")
    _password := c.FormValue("password")
    if auth == "" || _password == "" {
        return c.JSON(400, map[string]string{"error": "Missing required fields"})
    }

    username, err := repository.Login(auth, _password) 
    if err != nil {
        fmt.Println(err)
        return c.JSON(401, map[string]string{"error": "Invalid credentials"})
    }

    token, err := encrypt.GenerateJWT(username)
    if err != nil {
        fmt.Println(err)
        return c.JSON(500, map[string]string{"error": "Internal Server Error"})
    }
    return c.JSON(200, map[string]string{"token": string(token)})
}