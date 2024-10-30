package environment

import (
	"os"
	"github.com/joho/godotenv"
)

type environment struct {
    GeminiToken string
    DatabaseURL string
    JWTSecret string
}

func load() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
}

func Get() (environment, error) {
    load()
    return environment{
        GeminiToken:   os.Getenv("GEMINI_TOKEN"),
        DatabaseURL: os.Getenv("DATABASE_URL"),
        JWTSecret: os.Getenv("JWT_SECRET"),
    }, nil
}