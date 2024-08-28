package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}
	return port
}
