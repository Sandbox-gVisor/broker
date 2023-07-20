package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	loadedConfig := Config{}

	loadedConfig.Address = os.Getenv("ADDRESS")
	loadedConfig.Type = os.Getenv("TYPE")

	if loadedConfig.Address == "" {
		return Config{}, errors.New("cant load .env")
	}
	return loadedConfig, nil
}
