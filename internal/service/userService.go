package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {

}

func (s UserService) Register(input dto.UserSignup) (string, error) {
	log.Println(input)

	return "demo-token-001", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	return nil, nil
}

func (s UserService) Login(input any) (string, error) {

	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]any, error) {

	return nil, nil
}

func (s UserService) CreateCart(id uint, u domain.User) ([]any, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]any, error) {

	return nil, nil
}

func (s UserService) GetOrderByID(id, userID uint) (any, error) {

	return nil, nil
}

