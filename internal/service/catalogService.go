package service

import (
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
