package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func initTransactionService(db *gorm.DB, auth helper.Auth) service.TransactionService {
	return service.TransactionService{
		Repo: repository.NewTransactionRepository(db),
		Auth: auth,
	}
}

func SetupTransactionRoutes(as *rest.RestHandler) {
	app := as.App
	svc := initTransactionService(as.DB, as.Auth)

	handler := TransactionHandler{svc: svc}

	secRoutes := app.Group("/", as.Auth.Authorize)
	secRoutes.Get("/payment", handler.MakePayment)

	sellerRoutes := app.Group("/seller", as.Auth.AuthorizeSeller)
	sellerRoutes.Get("/orders", handler.GetOrders)
	sellerRoutes.Get("/orders/:id", handler.GetOrderDetails)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "payment", nil)
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "get order", nil)
}

func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "order details", nil)
}