package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvGet(key string, fallback string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
