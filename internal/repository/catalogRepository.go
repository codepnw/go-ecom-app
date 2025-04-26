package repository

import (
	"go-ecommerce-app/internal/domain"

	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

func (c catalogRepository) CreateCategory(e *domain.Category) error {
	if err := c.db.Create(&e).Error; err != nil {
		return err
	}
	return nil
}

func (c catalogRepository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category

	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c catalogRepository) FindCategoryByID(id int) (*domain.Category, error) {
	var category *domain.Category

	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (c catalogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	if err := c.db.Save(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (c catalogRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&domain.Category{}, id).Error
	return err
}
