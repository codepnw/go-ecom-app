package repository

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePayment(payment *domain.Payment) error
	FindInitialPayment(userID uint) (*domain.Payment, error)
	FindOrders(userID uint) ([]domain.OrderItem, error)
	FindOrderByID(userID, orderID uint) (dto.SellerOrderDetails, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreatePayment(payment *domain.Payment) error {
	return r.db.Create(payment).Error
}

func (r *transactionRepository) FindInitialPayment(userID uint) (*domain.Payment, error) {
	var payment *domain.Payment
	err := r.db.First(&payment, "user_id=? AND status=initial", userID).Order("created_at desc").Error
	return payment, err
}

func (r *transactionRepository) FindOrders(userID uint) ([]domain.OrderItem, error) {
	panic("")
}

func (r *transactionRepository) FindOrderByID(userID, orderID uint) (dto.SellerOrderDetails, error) {
	panic("")
}

