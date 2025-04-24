package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/notification"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) Register(input dto.UserSignup) (string, error) {
	hashedPassword, err := s.Auth.GenerateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)`:
			return "", errors.New("email already exists")
		default:
			return "", err
		}
	}

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) Login(email, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// vertify password
	if err = s.Auth.VerifyPassword(password, user.Password); err != nil {
		return "", err
	}

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserByID(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) error {
	// check user
	if s.isVerifiedUser(e.ID) {
		return errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return err
	}

	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	// update user
	if _, err = s.Repo.UpdateUser(e.ID, user); err != nil {
		return errors.New("unable to update verification code")
	}

	user, _ = s.Repo.FindUserByID(e.ID)

	msg := fmt.Sprintf("Your verification code is %s", code)

	notiClient := notification.NewNotificationClient(s.Config)
	// send SMS
	if err = notiClient.SendSMS(user.Phone, msg); err != nil {
		return errors.New("error on sending sms")
	}

	return nil
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
