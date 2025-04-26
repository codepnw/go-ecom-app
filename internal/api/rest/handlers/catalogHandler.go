package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := CatalogHandler{
		svc: svc,
	}

	// Public routes
	app.Get("/products")
	app.Get("/products/:id")
	app.Get("/categories")
	app.Get("/categories/:id")

	// Private routes
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	// categories
	selRoutes.Post("/categories", handler.CreateCategory)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Patch("/products/:id", handler.UpdateProduct) // update stock
	selRoutes.Put("/products/:id", handler.EditProduct)
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

// Categories
func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

// Products
func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}

func (h CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	return rest.SuccessResponse(ctx, "products endpoint", nil)
}
