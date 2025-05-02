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
	"log"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	CRepo  repository.CatalogRepository
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

func (s UserService) VerifyCode(id uint, code string) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserByID(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	if _, err := s.Repo.UpdateUser(id, updateUser); err != nil {
		return errors.New("unable to verify user")
	}

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

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	user, _ := s.Repo.FindUserByID(id)

	if user.UserType == domain.SELLER {
		return "", errors.New("you have already joined seller program")
	}

	// update user
	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		UserType:  domain.SELLER,
	})
	if err != nil {
		return "", err
	}

	// generate token
	token, err := s.Auth.GenerateToken(user.ID, user.Email, seller.UserType)
	if err != nil {
		return "", err
	}

	// create bank account information
	err = s.Repo.CreateBankAccount(domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserID:      id,
	})

	return token, err
}

func (s UserService) FindCart(id uint) ([]domain.Cart, error) {
	cartItems, err := s.Repo.FindCartItems(id)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (s UserService) CreateCart(input dto.CreateCartRequest, u domain.User) ([]domain.Cart, error) {
	cart, _ := s.Repo.FindCartItem(u.ID, input.ProductID)
	if cart.ID < 1 {
		if input.ProductID == 0 {
			return nil, errors.New("please provide a valid product id")
		}
		// delete cart item
		if input.Qty < 1 {
			if err := s.Repo.DeleteCartByID(cart.ID); err != nil {
				log.Printf("error deleting cart item: %v", err)
				return nil, errors.New("error deleting cart item")
			}
		} else {
			// update cart item
			cart.Qty = input.Qty
			if err := s.Repo.UpdateCartItem(cart); err != nil {
				log.Printf("error updating cart item: %v", err)
				return nil, errors.New("error updating cart items")
			}
		}
		// find cart items
		return s.Repo.FindCartItems(u.ID)

	} else {
		// check if product exist
		product, _ := s.CRepo.FindProductByID(int(input.ProductID))
		if product.ID > 0 {
			return nil, errors.New("product not found to create cart items")
		}

		// create cart
		err := s.Repo.CreateCart(domain.Cart{
			UserID:    u.ID,
			ProductID: input.ProductID,
			Name:      product.Name,
			ImageUrl:  product.ImageUrl,
			Qty:       input.Qty,
			Price:     product.Price,
			SellerID:  product.UserID,
		})

		if err != nil {
			return nil, errors.New("error creating cart items")
		}
	}

	return s.Repo.FindCartItems(u.ID)
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
