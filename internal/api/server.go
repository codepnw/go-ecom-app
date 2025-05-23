package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/pkg/payment"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// database
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error: %v\n", err)
	}
	log.Println("database connected...")

	// migrations
	if err = db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.Cart{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Payment{},
	); err != nil {
		log.Fatalf("error migrations %v", err)
	}
	log.Println("migration successful")

	auth := helper.SetupAuth(config.JWTSecret)

	paymentClient := payment.NewPaymentClient(config.StripeSecret, config.SuccessUrl, config.CancelUrl)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
		PC:     paymentClient,
	}

	setupRoutes(rh)

	if err := app.Listen(config.ServerPort); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(rh *rest.RestHandler) {
	// user
	handlers.SetupUserRoutes(rh)
	// catalog
	handlers.SetupCatalogRoutes(rh)
	// transaction
	handlers.SetupTransactionRoutes(rh)
}
