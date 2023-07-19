package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadedConfig := Config{}

	loadedConfig.Host = os.Getenv("HOST")
	port := os.Getenv("PORT")
	loadedConfig.Port = port

	if port == "" {
		// return default config
		loadedConfig.Host = "localhost"
		loadedConfig.Port = "9988"
		loadedConfig.Type = "tcp"
		return loadedConfig
	}
	loadedConfig.Type = os.Getenv("TYPE")
	return loadedConfig
}
