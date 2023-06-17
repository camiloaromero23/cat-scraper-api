package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
  err := godotenv.Load(".env")
	if err != nil {
    log.Println("No .env file found")
	}
  variable := os.Getenv(key)

	return variable
}
