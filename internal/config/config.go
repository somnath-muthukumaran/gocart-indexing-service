package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	RMQ_URL     string
	ES_URL      string
	ES_USERNAME string
	ES_PASSWORD string
}

var envConfig *Config

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Environment variable %s not set, using default value: %s", key, fallback)
	return fallback
}

func LoadConfig() {
	rootDir := findProjectRoot()
	envFile := filepath.Join(rootDir, ".env")
	fmt.Println("envFile path ", envFile)
	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env file found or error loading it, falling back to environment variables")
	}
	envConfig = &Config{
		RMQ_URL:     getEnv("RMQ_URL", "amqp://guest:guest@localhost:5672/"),
		ES_URL:      getEnv("ES_URL", "https://localhost:9200"),
		ES_USERNAME: getEnv("ES_USERNAME", "admin"),
		ES_PASSWORD: getEnv("ES_PASSWORD", "admin"),
	}
}

func findProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory:", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {

			return ""
		}
		currentDir = parentDir
	}
}

func GetEnvDetails() *Config {
	return envConfig
}
