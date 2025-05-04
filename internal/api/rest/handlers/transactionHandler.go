package handlers

import (
	"errors"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/pkg/payment"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	svc           service.TransactionService
	userSvc       service.UserService
	paymentClient payment.PaymentClient
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
	userSvc := service.UserService{
		Repo:   repository.NewUserRepository(as.DB),
		CRepo:  repository.NewCatalogRepository(as.DB),
		Auth:   as.Auth,
		Config: as.Config,
	}

	handler := TransactionHandler{
		svc:           svc,
		userSvc:       userSvc,
		paymentClient: as.PC,
	}

	secRoutes := app.Group("/", as.Auth.Authorize)
	secRoutes.Get("/payment", handler.MakePayment)

	sellerRoutes := app.Group("/seller", as.Auth.AuthorizeSeller)
	sellerRoutes.Get("/orders", handler.GetOrders)
	sellerRoutes.Get("/orders/:id", handler.GetOrderDetails)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	// check if payment session active
	activePayment, _ := h.svc.GetActivePayment(user.ID)
	if activePayment.ID > 0 {
		return rest.SuccessCreated(ctx, "create payment", activePayment.PaymentUrl)
	}

	// get total amount
	_, amount, err := h.userSvc.FindCart(user.ID)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	orderID, err := helper.RandomNumbers(8)
	if err != nil {
		return rest.InternalError(ctx, errors.New("error generating order id"))
	}

	// create a new payment session on stripe
	result, err := h.paymentClient.CreatePayment(amount, user.ID, orderID)
	if err != nil {
		return rest.BadRequestResponse(ctx, err.Error())
	}

	// create a new payment session to database
	err = h.svc.StoreCreatePayment(user.ID, result, amount, orderID)
	if err != nil {
		return rest.BadRequestResponse(ctx, err.Error())
	}

	return rest.SuccessCreated(ctx, "create payment", &fiber.Map{
		"result":      result,
		"payment_url": result.URL,
	})
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "get order", nil)
}

func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "order details", nil)
}
