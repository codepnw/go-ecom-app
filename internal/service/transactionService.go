package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
	Auth helper.Auth
}

func NewTransactionService(repo repository.TransactionRepository, auth helper.Auth) *TransactionService {
	return &TransactionService{
		Repo: repo,
		Auth: auth,
	}
}

func (s TransactionService) GetOrders(u domain.User) ([]domain.OrderItem, error) {
	orders, err := s.Repo.FindOrders(u.ID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s TransactionService) GetOrderDetails(u domain.User, orderID uint) (dto.SellerOrderDetails, error) {
	order, err := s.Repo.FindOrderByID(u.ID, orderID)
	if err != nil {
		return dto.SellerOrderDetails{}, err
	}

	return order, nil
}