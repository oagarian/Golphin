package encrypt

import (
	"golphin/src/utils/environment"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func IsPasswordValid(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func GenerateJWT(username string) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	secret, err := environment.Get()
	if err != nil {
        return nil, err
    }
	tokenString, err := token.SignedString([]byte(secret.JWTSecret))
	if err != nil {
		return nil, err
	}
	return []byte(tokenString), nil
}
