package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

const (
	errorNotFound = "record not found"
	notFoundResponse = "category id not found"
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
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProduct)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryByID)

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
func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	categories, err := h.svc.GetCategories()
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "success", categories)
}

func (h CatalogHandler) GetCategoryByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	category, err := h.svc.GetCategory(id)
	if err != nil {
		switch {
		case err.Error() == errorNotFound:
			return rest.NotFoundResponse(ctx, notFoundResponse)
		default:
			return rest.InternalError(ctx, err)
		}
	}

	return rest.SuccessResponse(ctx, "success", category)
}

func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestResponse(ctx, "create category is not valid")
	}

	// create category
	if err := h.svc.CreateCategory(req); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessCreated(ctx, "success", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	req := dto.CreateCategoryRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestResponse(ctx, "update category is not valid")
	}

	// update category
	category, err := h.svc.EditCategory(id, req)
	if err != nil {
		switch {
		case err.Error() == errorNotFound:
			return rest.NotFoundResponse(ctx, notFoundResponse)
		default:
			return rest.InternalError(ctx, err)
		}
	}

	return rest.SuccessResponse(ctx, "success", category)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.svc.DeleteCategory(id); err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.NoContentResponse(ctx)
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
