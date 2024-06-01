package elo

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Func to read API key from .env file
func GetAPIKey() (apiKey string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	result := os.Getenv("API_KEY")
	apiKey = result

	return
}
