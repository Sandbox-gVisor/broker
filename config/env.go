package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func checkConfig(config Config) error {
	if config.Address == "" {
		errors.New("cant load .env")
	} else if config.Type == "" {
		errors.New("cant load .env")
	} else if config.WebsoketPort == "" {
		errors.New("cant load .env")
	} else if config.RabbitAddress == "" {
		errors.New("cant load .env")
	} else if config.QueueName == "" {
		errors.New("cant load .env")
	}
	return nil
}

func loadEnv() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	loadedConfig := Config{}

	loadedConfig.Address = os.Getenv("ADDRESS")
	loadedConfig.Type = os.Getenv("TYPE")
	loadedConfig.WebsoketPort = os.Getenv("WS_PORT")
	loadedConfig.RabbitAddress = os.Getenv("RABBIT_ADDRESS")
	loadedConfig.QueueName = os.Getenv("QUEUE")

	err = checkConfig(loadedConfig)

	return loadedConfig, err
}
