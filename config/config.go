package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetJWTSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not found")
	}
	return secret
}
