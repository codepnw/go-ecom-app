package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv(envFileName string) (cfg AppConfig, err error) {
	if err := godotenv.Load(envFileName); err != nil {
		return AppConfig{}, errors.New("failed loading env file")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
