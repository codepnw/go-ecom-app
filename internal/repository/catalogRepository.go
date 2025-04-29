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

	CreateProduct(e *domain.Product) error
	FindProducts() ([]*domain.Product, error)
	FindProductByID(id int) (*domain.Product, error)
	FindSellerProducts(id int) ([]*domain.Product, error)
	EditProduct(e *domain.Product) (*domain.Product, error)
	DeleteProduct(e *domain.Product) error
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

// Products
func (c *catalogRepository) CreateProduct(e *domain.Product) error {
	err := c.db.Model(&domain.Product{}).Create(e).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *catalogRepository) FindProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	if err := c.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (c *catalogRepository) FindProductByID(id int) (*domain.Product, error) {
	var product *domain.Product

	if err := c.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (c *catalogRepository) FindSellerProducts(id int) ([]*domain.Product, error) {
	var products []*domain.Product

	err := c.db.Where("user_id=?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c *catalogRepository) EditProduct(e *domain.Product) (*domain.Product, error) {
	if err := c.db.Save(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (c *catalogRepository) DeleteProduct(e *domain.Product) error {
	if err := c.db.Delete(&domain.Product{}, e.ID).Error; err != nil {
		return err
	}
	return nil
}
