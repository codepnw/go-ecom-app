package repository

import (
	"go-ecommerce-app/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error

	// Cart
	FindCartItems(userID uint) ([]domain.Cart, error)
	FindCartItem(userID, productID uint) (domain.Cart, error)
	CreateCart(c domain.Cart) error
	UpdateCartItem(c domain.Cart) error
	DeleteCartByID(id uint) error
	DeleteCartItems(userID uint) error

	// Profile
	CreateProfile(e domain.Address) error
	UpdateProfile(e domain.Address) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.Preload("Address").First(&user, "email=?", email).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User

	if err := r.db.Preload("Address").First(&user, id).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(&u).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

// Cart
func (r userRepository) FindCartItems(userID uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id=?", userID).Find(&carts).Error

	return carts, err
}

func (r userRepository) FindCartItem(userID, productID uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id=? AND product_id=?", userID, productID).First(&cartItem).Error

	return cartItem, err
}

func (r userRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

func (r userRepository) UpdateCartItem(c domain.Cart) error {
	var cart domain.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.ID).Updates(c).Error

	return err
}

func (r userRepository) DeleteCartByID(id uint) error {
	return r.db.Delete(&domain.Cart{}, id).Error
}

func (r userRepository) DeleteCartItems(userID uint) error {
	err := r.db.Where("user_id=?", userID).Delete(&domain.Cart{}).Error
	return err
}

// Profile
func (r userRepository) CreateProfile(e domain.Address) error {
	return r.db.Create(&e).Error
}

func (r userRepository) UpdateProfile(e domain.Address) error {
	err := r.db.Where("user_id=?", e.UserID).Updates(e).Error
	return err
}
