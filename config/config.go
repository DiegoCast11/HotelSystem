package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetJWTSecret() string {
	if os.Getenv("ENV") != "production" { // Asume que ENV se configura como "production" en Heroku
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error al cargar el archivo .env")
		}
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not found")
	}
	return secret
}
