package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	DSN        string
	JWTSecret  string
}

func SetupEnv(envFileName string) (cfg AppConfig, err error) {
	if err := godotenv.Load(envFileName); err != nil {
		return AppConfig{}, errors.New("failed loading env file")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("http port variable not found")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("dsn variable not found")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("jwt secret variable not found")
	}

	return AppConfig{
		ServerPort: httpPort,
		DSN:        dsn,
		JWTSecret:  jwtSecret,
	}, nil
}
