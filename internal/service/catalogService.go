package service

import (
	"errors"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"strconv"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatalogService) GetCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, err
	}

	return categories, err
}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	category, err := s.Repo.FindCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, err
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {
	category, err := s.Repo.FindCategoryByID(id)
	if err != nil {
		return nil, err
	}

	if len(input.Name) > 0 {
		category.Name = input.Name
	}

	if input.ParentID > 0 {
		parentID := strconv.Itoa(input.ParentID)
		category.ParentID = parentID
	}

	if len(input.ImageUrl) > 0 {
		category.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		category.DisplayOrder = input.DisplayOrder
	}

	updated, err := s.Repo.EditCategory(category)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s CatalogService) DeleteCategory(id int) error {
	return s.Repo.DeleteCategory(id)
}

// Products

func (s CatalogService) CreateProduct(input dto.CreateProductRequest, user domain.User) error {
	err := s.Repo.CreateProduct(&domain.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		ImageUrl:    input.ImageUrl,
		UserID:      user.ID,
		Stock:       uint(input.Stock),
	})
	return err
}

func (s CatalogService) EditProduct(id int, input dto.CreateProductRequest, user domain.User) (*domain.Product, error) {
	product, err := s.Repo.FindProductByID(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	// verify owner
	if product.UserID != user.ID {
		return nil, errors.New("you dont have manage rights of this product")
	}

	if len(input.Name) > 0 {
		product.Name = input.Name
	}

	if len(input.Description) > 0 {
		product.Description = input.Description
	}

	if input.Price > 0 {
		product.Price = input.Price
	}

	if input.CategoryID > 0 {
		product.CategoryID = input.CategoryID
	}

	return s.Repo.EditProduct(product)
}

func (s CatalogService) DeleteProduct(id int, user domain.User) error {
	product, err := s.Repo.FindProductByID(id)
	if err != nil {
		return errors.New("product does not exist")
	}

	if product.UserID != user.ID {
		return errors.New("you dont have manage right of product")
	}

	if err = s.Repo.DeleteProduct(product); err != nil {
		return errors.New("product cant delete")
	}

	return nil
}

func (s CatalogService) GetProducts() ([]*domain.Product, error) {
	products, err := s.Repo.FindProducts()
	if err != nil {
		return nil, errors.New("products does not exist")
	}

	return products, nil
}

func (s CatalogService) GetProductByID(id int) (*domain.Product, error) {
	product, err := s.Repo.FindProductByID(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	return product, nil
}

func (s CatalogService) GetSellerProducts(id int) ([]*domain.Product, error) {
	products, err := s.Repo.FindSellerProducts(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	return products, nil
}

func (s CatalogService) UpdateProductStock(e domain.Product) (*domain.Product, error) {
	product, err := s.Repo.FindProductByID(int(e.ID))
	if err != nil {
		return nil, errors.New("product not found")
	}	

	// verify owner
	if product.UserID != e.UserID {
		return nil, errors.New("you dont have manage right of product")
	}

	product.Stock = e.Stock 
	editProduct, err := s.Repo.EditProduct(product)
	if err != nil {
		return nil, err
	}

	return editProduct, nil
}