package main

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"log"
)

const envFileName = "dev.env"

func main() {
	cfg, err := config.SetupEnv(envFileName)
	if err != nil {
		log.Fatalf("config setup failed: %v", err)
	}

	api.StartServer(cfg)
}
